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

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

type Text struct {
	Face  font.Face
	Rect  Rect
	Align Align
}

type Typer struct {
	Face  font.Face
	Rect  Rect
	Align Align
}

func NewTyper(tt *truetype.Font, base Rect, spacing float64, align Align) *Typer {
	t := &Typer{}
	r := base
	r.Bottom = r.Top + r.Height()/spacing
	t.Rect = r
	t.Face = truetype.NewFace(tt, &truetype.Options{
		Size:    t.Rect.Height(),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	t.Align = align
	return t
}

func (t *Typer) Draw(image *et.Image, str string, clr color.Color) {
	var x, y float64
	switch t.Align {
	case AlignLeft:
		x, y = t.Rect.AnchorPos(7)
	case AlignCenter:
		x, y = t.Rect.AnchorPos(8)
	case AlignRight:
		x, y = t.Rect.AnchorPos(9)
	}
	text.Draw(image, str, t.Face, int(x+0.5), int(y+0.5), clr)
}
