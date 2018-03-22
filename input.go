package tendon

import (
	et "github.com/hajimehoshi/ebiten"
)

var (
	Pressed, JustPressed bool
	Cursor               complex128
)

func updateInput() {
	x, y := et.CursorPosition()
	p := et.IsMouseButtonPressed(et.MouseButtonLeft)
	for _, t := range et.Touches() {
		x, y = t.Position()
		if x+y > 0 {
			pressed = true
		}
	}
	Cursor = complex(float64(x), float64(y))

	if p {
		if !Pressed {
			JustPressed = true
		} else {
			JustPressed = false
		}
		Pressed = true
	} else {
		Pressed = false
		JustPressed = false
	}
}
