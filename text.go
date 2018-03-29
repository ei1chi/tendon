// see github.com/explodes/tempura/blob/master/text.go
package tendon

import (
	"fmt"
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

const dpi = 72

func NewFont(path string) (*truetype.Font, error) {
	f, err := ebitenutil.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	tt, err := truetype.Parse(b)
	if err != nil {
		return nil, err
	}

	return tt, err
}

func NewFontFace(tt *truetype.Font, size float64) font.Face {
	return truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func MeasureText(text string, face font.Face) (int, int, int) {
	bounds, a := font.BoundString(face, text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	h := (bounds.Max.Y - bounds.Min.Y).Ceil()
	advance := a.Ceil()
	return w, h, advance
}

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

type Text struct {
	Face    font.Face
	Color   color.Color
	Text    string
	W, H    int
	Advance int
}

func NewText(face font.Face, color color.Color, text string) *Text {
	w, h, advance := MeasureText(text, face)
	return &Text{
		Face:    face,
		Color:   color,
		Text:    text,
		W:       w,
		H:       h,
		Advance: advance,
	}
}

func NewTextf(face font.Face, color color.Color, format string, args ...interface{}) *Text {
	if len(args) == 0 {
		return NewText(face, color, format)
	}
	return NewText(face, color, fmt.Sprintf(format, args...))
}

func (t *Text) Draw(image *et.Image, x, y float64, align Align) {
	switch align {
	case AlignCenter:
		x = x - float64(t.W)/2
	case AlignRight:
		x = x - float64(t.W)
	}
	text.Draw(image, t.Text, t.Face, int(x), int(y), t.Color)
}
