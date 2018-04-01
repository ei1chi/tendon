package tendon

import (
	et "github.com/hajimehoshi/ebiten"
)

var (
	IsPressed, IsJustPressed, IsJustReleased bool
	CursorPos                                complex128
)

func UpdateInput() {
	x, y := et.CursorPosition()
	p := et.IsMouseButtonPressed(et.MouseButtonLeft)
	for _, t := range et.Touches() {
		x, y = t.Position()
		if x+y > 0 {
			IsPressed = true
		}
	}
	CursorPos = complex(float64(x), float64(y))

	IsJustPressed = false
	IsJustReleased = false
	if p {
		if !IsPressed {
			IsJustPressed = true
		}
		IsPressed = true
	} else {
		if IsPressed {
			IsJustReleased = true
		}
		IsPressed = false
	}
}
