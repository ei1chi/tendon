package tendon

import (
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Sprite struct {
	Image    *et.Image
	halfSize complex128
}

func NewSprite(path string) (*Sprite, error) {
	s := &Sprite{}
	var err error
	s.Image, _, err = ebitenutil.NewImageFromFile(path, et.FilterDefault)
	w, h := s.Image.Size()
	s.halfSize = complex(float64(w)/2, float64(h)/2)
	return s, err
}

func (s *Sprite) Center() *et.DrawImageOptions {
	op := &et.DrawImageOptions{}
	op.GeoM.Translate(-real(s.halfSize), -imag(s.halfSize))
	return op
}
