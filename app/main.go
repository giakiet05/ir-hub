package main

import (
	"app/internal/config"
	"app/internal/handler"
	"app/internal/logging"
	"app/internal/serial"
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bendahl/uinput"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// Set up signal handling to gracefully shut down the application
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	byIdPath := os.Getenv("BY_ID_PATH")
	baudRateStr := os.Getenv("BAUD_RATE")
	baudRate, _ := strconv.Atoi(baudRateStr)

	presetDir := os.Getenv("PRESET_DIR")
	if presetDir == "" {
		presetDir = "presets"
	}

	if byIdPath == "" || baudRate == 0 {
		panic("BY_ID_PATH and BAUD_RATE must be set in the environment")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("ir-hub-keyboard"))
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	mouse, err := uinput.CreateMouse("/dev/uinput", []byte("ir-hub-mouse"))
	if err != nil {
		panic(err)
	}
	defer mouse.Close()

	logger := logging.NewLogger()

	irActionContext := handler.NewIrActionContext(keyboard, mouse, logger)
	irHandler := handler.NewIrEventHandler(irActionContext, logger)

	handler.RegisterDefaultActions(irHandler)

	newKeyActionMap, err := handler.ParsePresetDir(presetDir)
	if err != nil {
		logger.Error("Failed to parse initial preset directory", "error", err, "path", presetDir)
	} else {
		irHandler.ReloadKeyActions(newKeyActionMap)
	}

	config.WatchPresetDir(presetDir, irHandler, logger)

	serial.StartTransport(ctx, byIdPath, baudRate, irHandler.Handle, logger)

	// Wait for a termination signal to gracefully shut down the application
	sig := <-sigChan
	logger.Info("Received signal, Shutting down...", "signal", sig)
	cancel()
}
