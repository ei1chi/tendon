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

func NewText(face font.Face, anchor int, str string) *Text {
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

func (t *Text) Draw(image *et.Image, x, y float64, clr color.Color) {
	w, h := t.Bounds.AnchorPos(t.Anchor)
	x -= w
	y += t.Bounds.Bottom - h
	text.Draw(image, t.Str, t.Face, int(x+0.5), int(y+0.5), clr)
}

type TextBox struct {
	T *Text
	R Rect
}

func NewTextBox(r Rect, face font.Face, anchor int, str string) *TextBox {
	t := &TextBox{}
	t.T = NewText(face, anchor, str)
	t.R = r
	return t
}

func (t *TextBox) Draw(image *et.Image, ofsx, ofsy float64, clr color.Color) {
	x, y := t.R.AnchorPos(t.T.Anchor)
	x += ofsx
	y += ofsy
	t.T.Draw(image, x, y, clr)
}

func (t *TextBox) Fit() *TextBox {
	w, h, _ := MeasureText(t.T.Str, t.T.Face)
	t.R = t.R.SnapInside(t.T.Anchor, float64(w), float64(h))
	return t
}
