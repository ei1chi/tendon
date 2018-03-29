package tendon

import (
	"sync"

	et "github.com/hajimehoshi/ebiten"
)

type Scene interface {
	Load()
	Update(screen *et.Image) (Scene, error)
}

func StartLoading(s Scene, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(s Scene) {
		s.Load()
		wg.Done()
	}(s)
}
