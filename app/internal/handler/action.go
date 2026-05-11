package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os/exec"
	"time"

	"github.com/bendahl/uinput"
)

type IrActionId string

const (
	IrActionVolume IrActionId = "volume"
	IrActionMute   IrActionId = "mute"

	IrActionMouseClick  IrActionId = "mouse_click"
	IrActionMouseMove   IrActionId = "mouse_move"
	IrActionMouseScroll IrActionId = "mouse_scroll"

	IrActionRunCommand IrActionId = "run_command"

	IrActionPlayPause     IrActionId = "play_pause"
	IrActionNextTrack     IrActionId = "next_track"
	IrActionPreviousTrack IrActionId = "previous_track"

	IrActionKeyboardShortcut IrActionId = "keyboard_shortcut"
)

type IrActionContext struct {
	keyboard uinput.Keyboard
	mouse    uinput.Mouse
	logger   *slog.Logger
}

func NewIrActionContext(keyboard uinput.Keyboard, mouse uinput.Mouse, logger *slog.Logger) *IrActionContext {
	if keyboard == nil || mouse == nil || logger == nil {
		panic("All parameters for IrActionContext must be provided")
	}

	return &IrActionContext{
		keyboard: keyboard,
		mouse:    mouse,
		logger:   logger,
	}
}

type IrAction func(ctx *IrActionContext, rawParams json.RawMessage) error

func Volume(ctx *IrActionContext, rawParams json.RawMessage) error {
	var p VolumeParams

	if err := json.Unmarshal(rawParams, &p); err != nil {
		return errors.New("invalid parameters for volume action")
	}

	switch p.Delta {
	case -1:
		return ctx.keyboard.KeyPress(uinput.KeyVolumedown)
	case 1:
		return ctx.keyboard.KeyPress(uinput.KeyVolumeup)
	default:
		return errors.New("invalid parameter value for volume action")
	}

}

func Mute(ctx *IrActionContext, _ json.RawMessage) error {
	return ctx.keyboard.KeyPress(uinput.KeyMute)
}

func MouseClick(ctx *IrActionContext, rawParams json.RawMessage) error {
	var p MouseClickParams

	if err := json.Unmarshal(rawParams, &p); err != nil {
		return errors.New("invalid parameters for mouse click action")
	}

	var clickFunc func() error

	switch p.Button {
	case MouseButtonRight:
		clickFunc = ctx.mouse.RightClick
	case MouseButtonMiddle:
		clickFunc = ctx.mouse.MiddleClick
	case MouseButtonLeft:
		clickFunc = ctx.mouse.LeftClick
	default:
		return errors.New("invalid parameter value for mouse click action")
	}

	for i := 0; i < p.Count; i++ {
		if err := clickFunc(); err != nil {
			return err
		}
	}

	return nil
}

func MouseMove(ctx *IrActionContext, rawParams json.RawMessage) error {
	var p MouseMoveParams

	if err := json.Unmarshal(rawParams, &p); err != nil {
		return errors.New("invalid parameters for mouse move action")
	}

	return ctx.mouse.Move(int32(p.DX), int32(p.DY))
}

func MouseScroll(ctx *IrActionContext, rawParams json.RawMessage) error {
	var p MouseScrollParams

	if err := json.Unmarshal(rawParams, &p); err != nil {
		return errors.New("invalid parameters for mouse scroll action")
	}

	var direction bool
	switch p.Direction {
	case MouseScrollHorizontal:
		direction = true
	case MouseScrollVertical:
		direction = false
	default:
		return errors.New("invalid parameter value for mouse scroll action: direction must be 'horizontal' or 'vertical'")
	}

	return ctx.mouse.Wheel(direction, int32(p.Amount))
}

func PlayPause(ctx *IrActionContext, _ json.RawMessage) error {
	return ctx.keyboard.KeyPress(uinput.KeyPlaypause)
}

func NextTrack(ctx *IrActionContext, _ json.RawMessage) error {
	return ctx.keyboard.KeyPress(uinput.KeyNextsong)
}

func PreviousTrack(ctx *IrActionContext, _ json.RawMessage) error {
	return ctx.keyboard.KeyPress(uinput.KeyPrevioussong)
}

func RunCommand(ctx *IrActionContext, rawParams json.RawMessage) error {
	var p RunCommandParams

	if err := json.Unmarshal(rawParams, &p); err != nil {
		return errors.New("invalid parameters for run command action")
	}

	if p.Command == "" {
		return errors.New("command cannot be empty for run command action")
	}

	cmd := exec.Command(p.Command, p.Args...)

	err := cmd.Start()
	if err != nil {
		return errors.New("failed to start command")
	}

	go func() {
		err := cmd.Wait()
		if err != nil {
			ctx.logger.Error("Command execution failed", "command", p.Command, "args", p.Args, "error", err)
		}
		ctx.logger.Debug("Command execution finished", "command", p.Command, "args", p.Args)

	}()

	return nil
}

func KeyboardShortcut(ctx *IrActionContext, rawParams json.RawMessage) error {
	var p KeyboardShortcutParams

	if err := json.Unmarshal(rawParams, &p); err != nil {
		return errors.New("invalid parameters for keyboard shortcut action")
	}

	if len(p.Keys) == 0 {
		return errors.New("keys cannot be empty for keyboard shortcut action")
	}

	for _, key := range p.Keys {
		uinputKey, ok := keyboardKeyToUinputKey[key]
		if !ok {
			return errors.New("invalid key in keyboard shortcut action: " + string(key))
		}

		if err := ctx.keyboard.KeyDown(uinputKey); err != nil {
			return err
		}

		defer func() {
			if err := ctx.keyboard.KeyUp(uinputKey); err != nil {
				ctx.logger.Error("Failed to release key in keyboard shortcut action", "key", key, "error", err)
			}
		}()

	}

	time.Sleep(100 * time.Millisecond)
	return nil
}
