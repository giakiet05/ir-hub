package main

import (
	"app/internal/config"
	"app/internal/handler"
	"app/internal/logging"
	"app/internal/serial"
	"context"
	"os"
	"strconv"

	"github.com/bendahl/uinput"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	byIdPath := os.Getenv("BY_ID_PATH")
	baudRateStr := os.Getenv("BAUD_RATE")
	baudRate, err := strconv.Atoi(baudRateStr)

	if err != nil {
		panic("BAUD_RATE must be a valid integer")
	}

	if byIdPath == "" || baudRate == 0 {
		panic("BY_ID_PATH and BAUD_RATE must be set in the environment")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("desk-station"))
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	mouse, err := uinput.CreateMouse("/dev/uinput", []byte("desk-station-mouse"))
	if err != nil {
		panic(err)
	}
	defer mouse.Close()

	logger := logging.NewLogger()

	irActionContext := handler.NewIrActionContext(keyboard, mouse, logger)
	irHandler := handler.NewIrEventHandler(irActionContext, logger)

	handler.RegisterDefaultActions(irHandler)

	presetDir := "presets"
	newKeyActionMap, err := handler.ParsePresetDir(presetDir)
	if err != nil {
		panic(err)
	}
	irHandler.ReloadKeyActions(newKeyActionMap)

	config.WatchPresetDir(presetDir, irHandler, logger)

	serial.StartTransport(ctx, byIdPath, baudRate, irHandler.Handle, logger)

	select {}
}
