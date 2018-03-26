package tendon

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"path/filepath"

	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type SpriteSheet struct {
	Frames map[string]SpriteInfo `json:"frames"`
	Meta   interface{}           `json:"meta"`
}

type SpriteInfo struct {
	Frame struct {
		X int `json:"x"`
		Y int `json:"Y"`
		W int `json:"W"`
		H int `json:"H"`
	} `json:"frame"`
}

type Atlas struct {
	name  string
	image *et.Image
	sheet SpriteSheet
}

// NewAtlas creates Atlas which contains image and SpriteSheet. image is path.png, sheet is path.json.
func NewAtlas(path string) (*Atlas, error) {
	a := &Atlas{}
	a.name = filepath.Base(path)
	var err error
	a.image, _, err = ebitenutil.NewImageFromFile(path+".png", et.FilterDefault)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(path + ".json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &a.sheet)
	if err != nil {
		return nil, err
	}

	return a, nil
}

type Sprite struct {
	atlas *Atlas
	info  *SpriteInfo
}

func (a *Atlas) NewSprite(name string) (*Sprite, error) {
	s := &Sprite{}
	s.atlas = a
	i, ok := a.sheet.Frames[name]
	if !ok {
		return nil, fmt.Errorf("Sprite Name %s not found in Atlas %s", name, a.name)
	}
	s.info = &i
	return s, nil
}

func (s *Sprite) Draw(dst *et.Image, op *et.DrawImageOptions) {
	f := s.info.Frame
	min := image.Point{f.X, f.Y}
	max := min.Add(image.Point{f.W, f.H})
	op.SourceRect = &image.Rectangle{min, max}
	dst.DrawImage(s.atlas.image, op)
}

func (s *Sprite) Center() *et.DrawImageOptions {
	op := et.DrawImageOptions{}
	f := s.info.Frame
	op.GeoM.Translate(-float64(f.W)/2, -float64(f.H)/2)
	return &op
}
