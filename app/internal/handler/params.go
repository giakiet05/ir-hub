package handler

type MouseButton string

const (
	MouseButtonLeft   MouseButton = "left"
	MouseButtonRight  MouseButton = "right"
	MouseButtonMiddle MouseButton = "middle"
)

type MouseScrollDirection string

const (
	MouseScrollVertical   MouseScrollDirection = "vertical"
	MouseScrollHorizontal MouseScrollDirection = "horizontal"
)

type VolumeParams struct {
	Delta int `json:"delta"` // -1 for volume down, 1 for volume up or negative for multiple steps down, positive for multiple steps up
}

type MouseMoveParams struct {
	DX int `json:"dx"`
	DY int `json:"dy"`
}

type MouseClickParams struct {
	Button MouseButton `json:"button"` // -1: right, 0: middle, 1: left
	Count  int         `json:"count"`
}

type MouseScrollParams struct {
	Direction MouseScrollDirection `json:"direction"` // "horizontal" or "vertical"
	Amount    int                  `json:"amount"`    //positive for scroll up/right, negative for scroll down/left
}

type RunCommandParams struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}
type KeyboardShortcutParams struct {
	Keys []KeyboardKey `json:"keys"` // e.g. ["ctrl", "alt", "t"]
}
