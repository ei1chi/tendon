package tendon

import (
	et "github.com/hajimehoshi/ebiten"
)

var (
	IsPressed, IsJustPressed bool
	CursorPos                complex128
)

func updateInput() {
	x, y := et.CursorPosition()
	p := et.IsMouseButtonPressed(et.MouseButtonLeft)
	for _, t := range et.Touches() {
		x, y = t.Position()
		if x+y > 0 {
			IsPressed = true
		}
	}
	CursorPos = complex(float64(x), float64(y))

	if p {
		if !IsPressed {
			IsJustPressed = true
		} else {
			IsJustPressed = false
		}
		IsPressed = true
	} else {
		IsPressed = false
		IsJustPressed = false
	}
}
