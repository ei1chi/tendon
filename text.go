// see github.com/explodes/tempura/blob/master/text.go
package tendon

import (
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

type Text struct {
	Face   font.Face
	Str    string
	Bounds Rect
	Anchor int
}

func NewText(tt *truetype.Font, size float64, anchor int, str string) *Text {
	face := NewFontFace(tt, size)
	b, _ := font.BoundString(face, str)
	w := (b.Max.X - b.Min.X).Ceil()
	h := (-b.Min.Y).Ceil()
	r := Rect{
		Left:   0,
		Top:    0,
		Right:  float64(w),
		Bottom: float64(h),
	}
	return &Text{
		Face:   face,
		Str:    str,
		Bounds: r,
		Anchor: anchor,
	}
}

func (t *Text) SetText(str string) {
	b, _ := font.BoundString(t.Face, str)
	w := (b.Max.X - b.Min.X).Ceil()
	h := (-b.Min.Y).Ceil()
	r := Rect{
		Left:   0,
		Top:    0,
		Right:  float64(w),
		Bottom: float64(h),
	}
	t.Str = str
	t.Bounds = r
}

// DrawA draws text at Absolute position
func (t *Text) DrawA(image *et.Image, x, y float64, clr color.Color) {
	w, h := t.Bounds.AnchorPos(t.Anchor)
	x -= w
	y += t.Bounds.Bottom - h
	text.Draw(image, t.Str, t.Face, int(x+0.5), int(y+0.5), clr)
}

// DrawR draw text on the Rect
func (t *Text) DrawR(image *et.Image, rect Rect, clr color.Color) {
	x, y := rect.AnchorPos(t.Anchor)
	t.DrawA(image, x, y, clr)
}

type TextBox struct {
	T *Text
	R Rect
}

func NewTextBox(r Rect, tt *truetype.Font, size float64, anchor int, str string) *TextBox {
	t := &TextBox{}
	t.T = NewText(tt, size, anchor, str)
	t.R = r
	return t
}

func (t *TextBox) Draw(image *et.Image, clr color.Color) {
	t.T.DrawR(image, t.R, clr)
}
