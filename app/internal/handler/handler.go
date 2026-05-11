package handler

import (
	"app/internal/serial"
	"encoding/json"
	"errors"
	"log/slog"
	"sync"
	"time"
)

type BoundAction struct {
	ActionId     IrActionId
	Repeatable   bool
	HoldDuration time.Duration
	Params       json.RawMessage
}

type IrEventHandler struct {
	mu           sync.RWMutex
	ctx          *IrActionContext
	keyActionMap map[string]BoundAction
	actionMap    map[IrActionId]IrAction
	logger       *slog.Logger

	lastAddr      uint16
	lastCmd       uint16
	lastPressTime time.Time
	holdStartTime time.Time
	actionFired   bool
}

func NewIrEventHandler(ctx *IrActionContext, logger *slog.Logger) *IrEventHandler {
	if ctx == nil {
		panic("IrActionContext cannot be nil")

	}

	if ctx.keyboard == nil || ctx.mouse == nil {
		panic("IrActionContext must have both Keyboard and Mouse initialized")
	}

	if logger == nil {
		panic("Logger cannot be nil")
	}

	return &IrEventHandler{
		ctx:          ctx,
		keyActionMap: make(map[string]BoundAction),
		actionMap:    make(map[IrActionId]IrAction),
		logger:       logger,
	}
}

func (eh *IrEventHandler) RegisterAction(actionId IrActionId, action IrAction) {
	eh.actionMap[actionId] = action
	eh.logger.Debug("Action registered", "action_id", actionId)
}

func (eh *IrEventHandler) fireAction(boundAction BoundAction) {
	action, exists := eh.actionMap[boundAction.ActionId]
	if !exists {
		eh.logger.Error("Action not found", "action_id", boundAction.ActionId)
		return
	}

	go func() {
		if err := action(eh.ctx, boundAction.Params); err != nil {
			eh.logger.Error("Failed to execute action", "action_id", boundAction.ActionId, "error", err)
		} else {
			eh.logger.Debug("Action executed successfully", "action_id", boundAction.ActionId)
		}
	}()
}

func (eh *IrEventHandler) Handle(payload *serial.IrMessage) error {
	if payload == nil {
		return errors.New("payload is nil")
	}

	key := IrKey{
		Address: payload.Address,
		Command: payload.Command,
	}
	keyStr := key.String()

	eh.mu.RLock()
	boundAction, exists := eh.keyActionMap[keyStr]
	eh.mu.RUnlock()

	if !exists {
		eh.logger.Debug("No action mapped for IR key", "key", keyStr)
		return nil
	}

	eh.mu.Lock()
	defer eh.mu.Unlock()

	now := time.Now()

	// FLOW 1: Handle repeated signal
	if payload.IsRepeat {
		if eh.lastAddr != payload.Address || eh.lastCmd != payload.Command {
			return nil
		}

		if boundAction.Repeatable {
			// Fire action immediately if this button is repeatable
			eh.fireAction(boundAction)
		} else if boundAction.HoldDuration > 0 { // Only consider hold duration for non-repeatable buttons and if hold duration is set
			// For non-repeatable buttons, fire action if hold duration has passed since the first press
			if !eh.actionFired && now.Sub(eh.holdStartTime) >= boundAction.HoldDuration {
				eh.fireAction(boundAction)
				eh.actionFired = true
				eh.logger.Debug("Hold action fired on repeat signal", "key", keyStr)
			}
		}
		return nil
	}

	// FLOW 2: Handle new button press (is_repeat = false)

	// [Debounce] Chống dội phím: Bấm lại cùng 1 phím trong vòng 150ms -> Bỏ qua!
	if eh.lastAddr == payload.Address && eh.lastCmd == payload.Command && now.Sub(eh.lastPressTime) < 150*time.Millisecond {
		return nil
	}

	eh.lastAddr = payload.Address
	eh.lastCmd = payload.Command
	eh.lastPressTime = now
	eh.holdStartTime = now
	eh.actionFired = false

	if boundAction.HoldDuration == 0 {
		// If no hold duration is set, fire action immediately on press
		eh.fireAction(boundAction)
		eh.actionFired = true
		eh.logger.Debug("Action fired on button press", "key", keyStr)
	} else {
		eh.logger.Debug("Button pressed, waiting for hold duration", "key", keyStr, "hold_duration_ms", boundAction.HoldDuration.Milliseconds())
	}

	return nil
}

func (eh *IrEventHandler) ReloadKeyActions(newMap map[string]BoundAction) {
	eh.mu.Lock()
	defer eh.mu.Unlock()
	eh.keyActionMap = newMap
	eh.logger.Info("Key actions reloaded", "total_actions", len(newMap))
}
