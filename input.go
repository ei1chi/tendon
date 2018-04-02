package tendon

import (
	et "github.com/hajimehoshi/ebiten"
)

var (
	IsPressed, IsJustPressed, IsJustReleased bool
	CursorPos                                complex128
)

func UpdateInput() {

	// get mouse info
	x, y := et.CursorPosition()
	p := et.IsMouseButtonPressed(et.MouseButtonLeft)

	// get touch info
	for _, t := range et.Touches() {
		tx, ty := t.Position()
		if tx+ty > 0 {
			x, y = tx, ty
			p = true
		}
	}
	if x+y > 0 {
		CursorPos = complex(float64(x), float64(y))
	}

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
