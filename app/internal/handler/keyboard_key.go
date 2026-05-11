package handler

import "github.com/bendahl/uinput"

type KeyboardKey string

const (
	KeyA KeyboardKey = "a"
	KeyB KeyboardKey = "b"
	KeyC KeyboardKey = "c"
	KeyD KeyboardKey = "d"
	KeyE KeyboardKey = "e"
	KeyF KeyboardKey = "f"
	KeyG KeyboardKey = "g"
	KeyH KeyboardKey = "h"
	KeyI KeyboardKey = "i"
	KeyJ KeyboardKey = "j"
	KeyK KeyboardKey = "k"
	KeyL KeyboardKey = "l"
	KeyM KeyboardKey = "m"
	KeyN KeyboardKey = "n"
	KeyO KeyboardKey = "o"
	KeyP KeyboardKey = "p"
	KeyQ KeyboardKey = "q"
	KeyR KeyboardKey = "r"
	KeyS KeyboardKey = "s"
	KeyT KeyboardKey = "t"
	KeyU KeyboardKey = "u"
	KeyV KeyboardKey = "v"
	KeyW KeyboardKey = "w"
	KeyX KeyboardKey = "x"
	KeyY KeyboardKey = "y"
	KeyZ KeyboardKey = "z"

	KeyCtrl  KeyboardKey = "ctrl"
	KeyAlt   KeyboardKey = "alt"
	KeyShift KeyboardKey = "shift"
	KeyMeta  KeyboardKey = "meta" // Windows key or Command key
	KeySpace KeyboardKey = "space"

	KeyUp    KeyboardKey = "up"
	KeyDown  KeyboardKey = "down"
	KeyLeft  KeyboardKey = "left"
	KeyRight KeyboardKey = "right"

	KeyEnter       KeyboardKey = "enter"
	KeyEscape      KeyboardKey = "escape"
	KeyPrintScreen KeyboardKey = "print_screen"
	KeyPageUp      KeyboardKey = "page_up"
	KeyPageDown    KeyboardKey = "page_down"
	KeyDelete      KeyboardKey = "delete"
	KeyBackspace   KeyboardKey = "backspace"
	KeyHome        KeyboardKey = "home"
	KeyEnd         KeyboardKey = "end"
	KeyTab         KeyboardKey = "tab"
)

var keyboardKeyToUinputKey = map[KeyboardKey]int{
	KeyA:           uinput.KeyA,
	KeyB:           uinput.KeyB,
	KeyC:           uinput.KeyC,
	KeyD:           uinput.KeyD,
	KeyE:           uinput.KeyE,
	KeyF:           uinput.KeyF,
	KeyG:           uinput.KeyG,
	KeyH:           uinput.KeyH,
	KeyI:           uinput.KeyI,
	KeyJ:           uinput.KeyJ,
	KeyK:           uinput.KeyK,
	KeyL:           uinput.KeyL,
	KeyM:           uinput.KeyM,
	KeyN:           uinput.KeyN,
	KeyO:           uinput.KeyO,
	KeyP:           uinput.KeyP,
	KeyQ:           uinput.KeyQ,
	KeyR:           uinput.KeyR,
	KeyS:           uinput.KeyS,
	KeyT:           uinput.KeyT,
	KeyU:           uinput.KeyU,
	KeyV:           uinput.KeyV,
	KeyW:           uinput.KeyW,
	KeyX:           uinput.KeyX,
	KeyY:           uinput.KeyY,
	KeyZ:           uinput.KeyZ,
	KeyCtrl:        uinput.KeyLeftctrl,
	KeyAlt:         uinput.KeyLeftalt,
	KeyShift:       uinput.KeyLeftshift,
	KeyMeta:        uinput.KeyLeftmeta,
	KeyUp:          uinput.KeyUp,
	KeyDown:        uinput.KeyDown,
	KeyLeft:        uinput.KeyLeft,
	KeyRight:       uinput.KeyRight,
	KeyEnter:       uinput.KeyEnter,
	KeyEscape:      uinput.KeyEsc,
	KeyPrintScreen: uinput.KeyPrint,
	KeySpace:       uinput.KeySpace,
	KeyPageUp:      uinput.KeyPageup,
	KeyPageDown:    uinput.KeyPagedown,
	KeyDelete:      uinput.KeyDelete,
	KeyBackspace:   uinput.KeyBackspace,
	KeyHome:        uinput.KeyHome,
	KeyEnd:         uinput.KeyEnd,
	KeyTab:         uinput.KeyTab,
}
