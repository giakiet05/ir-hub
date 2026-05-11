package serial

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"time"

	se "github.com/tarm/serial"
)

func ParseIrMessage(line []byte) (IrMessage, error) {
	trimmed := bytes.TrimSpace(line)
	if len(trimmed) == 0 {
		return IrMessage{}, errors.New("empty message")
	}

	msg := IrMessage{}
	if err := json.Unmarshal(trimmed, &msg); err != nil {
		return IrMessage{}, err
	}

	return msg, nil
}

func StartTransport(ctx context.Context, byIdPath string, baudRate int, onMessage func(*IrMessage) error, logger *slog.Logger) {
	go func() {
		for {
			// Check if we should exit
			select {
			case <-ctx.Done():
				return
			default:
			}

			config := &se.Config{
				Name: byIdPath,
				Baud: baudRate,
			}
			port, err := se.OpenPort(config)
			if err != nil {
				logger.Warn("Failed to open serial port, retrying in 2s...", "error", err, "port", byIdPath)
				select {
				case <-ctx.Done():
					return
				case <-time.After(2 * time.Second):
					continue
				}
			}

			logger.Info("Serial transport connected", "port", byIdPath)
			runReadLoop(ctx, port, onMessage, logger)
			port.Close()

			if ctx.Err() != nil {
				return
			}
			logger.Warn("Serial connection lost, attempting to reconnect in 1s...")
			time.Sleep(1 * time.Second)
		}
	}()
}

func runReadLoop(ctx context.Context, port *se.Port, onMessage func(*IrMessage) error, logger *slog.Logger) {
	reader := bufio.NewReader(port)

	// Local context monitoring to force close port on shutdown
	go func() {
		<-ctx.Done()
		_ = port.Close()
	}()

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			logger.Error("Serial read error", "error", err)
			return
		}

		msg, err := ParseIrMessage(line)
		if err != nil {
			logger.Warn("Failed to parse IR message", "error", err, "line", string(line))
			continue
		}

		logger.Debug("IR message received", "raw_code", msg.RawCode)
		if onMessage != nil {
			if err := onMessage(&msg); err != nil {
				logger.Error("Failed to handle IR message", "error", err)
			}
		}
	}
}
