package sysimp

import rl "github.com/gen2brain/raylib-go/raylib"

type MouseButton int32

var GetKeyPressed = rl.GetKeyPressed
var IsKeyDown = rl.IsKeyDown
var IsKeyPressed = rl.IsKeyPressed
var IsKeyPressedRepeat = rl.IsKeyPressedRepeat
var IsKeyReleased = rl.IsKeyReleased
var IsKeyUp = rl.IsKeyUp
var SetExitKey = rl.SetExitKey

// Mouse Buttons
const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
	MouseButtonMiddle
	MouseButtonSide
	MouseButtonExtra
	MouseButtonForward
	MouseButtonBack
)

// Some basic Defines
const (
	// Raylib Config Flags
	// Set to try enabling V-Sync on GPU
	FlagVsyncHint = 0x00000040
	// Set to run program in fullscreen
	FlagFullscreenMode = 0x00000002
	// Set to allow resizable window
	FlagWindowResizable = 0x00000004
	// Set to disable window decoration (frame and buttons)
	FlagWindowUndecorated = 0x00000008
	// Set to hide window
	FlagWindowHidden = 0x00000080
	// Set to minimize window (iconify)
	FlagWindowMinimized = 0x00000200
	// Set to maximize window (expanded to monitor)
	FlagWindowMaximized = 0x00000400
	// Set to window non focused
	FlagWindowUnfocused = 0x00000800
	// Set to window always on top
	FlagWindowTopmost = 0x00001000
	// Set to allow windows running while minimized
	FlagWindowAlwaysRun = 0x00000100
	// Set to allow transparent window
	FlagWindowTransparent = 0x00000010
	// Set to support HighDPI
	FlagWindowHighdpi = 0x00002000
	// Set to support mouse passthrough, only supported when FLAG_WINDOW_UNDECORATED
	FlagWindowMousePassthrough = 0x00004000
	// Set to run program in borderless windowed mode
	FlagBorderlessWindowedMode = 0x00008000
	// Set to try enabling MSAA 4X
	FlagMsaa4xHint = 0x00000020
	// Set to try enabling interlaced video format (for V3D)
	FlagInterlacedHint = 0x00010000

	// KeyNull is used for no key pressed
	KeyNull = 0

	// Keyboard Function Keys
	KeySpace        = 32
	KeyEscape       = 256
	KeyEnter        = 257
	KeyTab          = 258
	KeyBackspace    = 259
	KeyInsert       = 260
	KeyDelete       = 261
	KeyRight        = 262
	KeyLeft         = 263
	KeyDown         = 264
	KeyUp           = 265
	KeyPageUp       = 266
	KeyPageDown     = 267
	KeyHome         = 268
	KeyEnd          = 269
	KeyCapsLock     = 280
	KeyScrollLock   = 281
	KeyNumLock      = 282
	KeyPrintScreen  = 283
	KeyPause        = 284
	KeyF1           = 290
	KeyF2           = 291
	KeyF3           = 292
	KeyF4           = 293
	KeyF5           = 294
	KeyF6           = 295
	KeyF7           = 296
	KeyF8           = 297
	KeyF9           = 298
	KeyF10          = 299
	KeyF11          = 300
	KeyF12          = 301
	KeyLeftShift    = 340
	KeyLeftControl  = 341
	KeyLeftAlt      = 342
	KeyLeftSuper    = 343
	KeyRightShift   = 344
	KeyRightControl = 345
	KeyRightAlt     = 346
	KeyRightSuper   = 347
	KeyKbMenu       = 348
	KeyLeftBracket  = 91
	KeyBackSlash    = 92
	KeyRightBracket = 93
	KeyGrave        = 96

	// Keyboard Number Pad Keys
	KeyKp0        = 320
	KeyKp1        = 321
	KeyKp2        = 322
	KeyKp3        = 323
	KeyKp4        = 324
	KeyKp5        = 325
	KeyKp6        = 326
	KeyKp7        = 327
	KeyKp8        = 328
	KeyKp9        = 329
	KeyKpDecimal  = 330
	KeyKpDivide   = 331
	KeyKpMultiply = 332
	KeyKpSubtract = 333
	KeyKpAdd      = 334
	KeyKpEnter    = 335
	KeyKpEqual    = 336

	// Keyboard Alpha Numeric Keys
	KeyApostrophe = 39
	KeyComma      = 44
	KeyMinus      = 45
	KeyPeriod     = 46
	KeySlash      = 47
	KeyZero       = 48
	KeyOne        = 49
	KeyTwo        = 50
	KeyThree      = 51
	KeyFour       = 52
	KeyFive       = 53
	KeySix        = 54
	KeySeven      = 55
	KeyEight      = 56
	KeyNine       = 57
	KeySemicolon  = 59
	KeyEqual      = 61
	KeyA          = 65
	KeyB          = 66
	KeyC          = 67
	KeyD          = 68
	KeyE          = 69
	KeyF          = 70
	KeyG          = 71
	KeyH          = 72
	KeyI          = 73
	KeyJ          = 74
	KeyK          = 75
	KeyL          = 76
	KeyM          = 77
	KeyN          = 78
	KeyO          = 79
	KeyP          = 80
	KeyQ          = 81
	KeyR          = 82
	KeyS          = 83
	KeyT          = 84
	KeyU          = 85
	KeyV          = 86
	KeyW          = 87
	KeyX          = 88
	KeyY          = 89
	KeyZ          = 90

	// Android keys
	KeyBack       = 4
	KeyMenu       = 5
	KeyVolumeUp   = 24
	KeyVolumeDown = 25

	MouseLeftButton   = MouseButtonLeft
	MouseRightButton  = MouseButtonRight
	MouseMiddleButton = MouseButtonMiddle
)

var keyToStringMap = map[int32]string{
	KeyNull:         "Null",
	KeySpace:        "Space",
	KeyEscape:       "Escape",
	KeyEnter:        "Enter",
	KeyTab:          "Tab",
	KeyBackspace:    "Backspace",
	KeyInsert:       "Insert",
	KeyDelete:       "Delete",
	KeyRight:        "Right",
	KeyLeft:         "Left",
	KeyDown:         "Down",
	KeyUp:           "Up",
	KeyPageUp:       "PageUp",
	KeyPageDown:     "PageDown",
	KeyHome:         "Home",
	KeyEnd:          "End",
	KeyCapsLock:     "CapsLock",
	KeyScrollLock:   "ScrollLock",
	KeyNumLock:      "NumLock",
	KeyPrintScreen:  "PrintScreen",
	KeyPause:        "Pause",
	KeyF1:           "F1",
	KeyF2:           "F2",
	KeyF3:           "F3",
	KeyF4:           "F4",
	KeyF5:           "F5",
	KeyF6:           "F6",
	KeyF7:           "F7",
	KeyF8:           "F8",
	KeyF9:           "F9",
	KeyF10:          "F10",
	KeyF11:          "F11",
	KeyF12:          "F12",
	KeyLeftShift:    "LeftShift",
	KeyLeftControl:  "LeftControl",
	KeyLeftAlt:      "LeftAlt",
	KeyLeftSuper:    "LeftSuper",
	KeyRightShift:   "RightShift",
	KeyRightControl: "RightControl",
	KeyRightAlt:     "RightAlt",
	KeyRightSuper:   "RightSuper",
	KeyKbMenu:       "KbMenu",
	KeyLeftBracket:  "LeftBracket",
	KeyBackSlash:    "BackSlash",
	KeyRightBracket: "RightBracket",
	KeyGrave:        "Grave",
	KeyKp0:          "Kp0",
	KeyKp1:          "Kp1",
	KeyKp2:          "Kp2",
	KeyKp3:          "Kp3",
	KeyKp4:          "Kp4",
	KeyKp5:          "Kp5",
	KeyKp6:          "Kp6",
	KeyKp7:          "Kp7",
	KeyKp8:          "Kp8",
	KeyKp9:          "Kp9",
	KeyKpDecimal:    "KpDecimal",
	KeyKpDivide:     "KpDivide",
	KeyKpMultiply:   "KpMultiply",
	KeyKpSubtract:   "KpSubtract",
	KeyKpAdd:        "KpAdd",
	KeyKpEnter:      "KpEnter",
	KeyKpEqual:      "KpEqual",
	KeyApostrophe:   "Apostrophe",
	KeyComma:        "Comma",
	KeyMinus:        "Minus",
	KeyPeriod:       "Period",
	KeySlash:        "Slash",
	KeyZero:         "Zero",
	KeyOne:          "One",
	KeyTwo:          "Two",
	KeyThree:        "Three",
	KeyFour:         "Four",
	KeyFive:         "Five",
	KeySix:          "Six",
	KeySeven:        "Seven",
	KeyEight:        "Eight",
	KeyNine:         "Nine",
	KeySemicolon:    "Semicolon",
	KeyEqual:        "Equal",
	KeyA:            "A",
	KeyB:            "B",
	KeyC:            "C",
	KeyD:            "D",
	KeyE:            "E",
	KeyF:            "F",
	KeyG:            "G",
	KeyH:            "H",
	KeyI:            "I",
	KeyJ:            "J",
	KeyK:            "K",
	KeyL:            "L",
	KeyM:            "M",
	KeyN:            "N",
	KeyO:            "O",
	KeyP:            "P",
	KeyQ:            "Q",
	KeyR:            "R",
	KeyS:            "S",
	KeyT:            "T",
	KeyU:            "U",
	KeyV:            "V",
	KeyW:            "W",
	KeyX:            "X",
	KeyY:            "Y",
	KeyZ:            "Z",
	KeyBack:         "Back",
	KeyMenu:         "Menu",
	KeyVolumeUp:     "VolumeUp",
	KeyVolumeDown:   "VolumeDown",
}

func KeyToString(key int32) string {
	return keyToStringMap[key]
}
