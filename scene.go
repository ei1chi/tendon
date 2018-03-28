package tendon

import (
	"errors"

	et "github.com/hajimehoshi/ebiten"
)

type Scene interface {
	Load()
	Update(*et.Image) (Scene, error)
}

type SceneBase struct {
	Slave     Scene
	task      chan struct{}
	hasLoaded bool
}

func (s *SceneBase) StartSlaveLoading() error {
	if s.task != nil {
		return errors.New("slave loading has already started!")
	}
	if s.Slave == nil {
		return errors.New("slave has not created!")
	}
	s.task = make(chan struct{}, 1)
	go func() {
		s.Slave.Load()
		<-s.task
	}()
	return nil
}

func (s *SceneBase) HasSlaveLoaded() (bool, error) {
	if s.task == nil {
		return false, errors.New("hasn't started loading")
	}
	if s.Slave == nil {
		return false, errors.New("slave has not created!")
	}
	if s.hasLoaded {
		return true, nil
	}
	select {
	case s.task <- struct{}{}: // 空である
		s.hasLoaded = true
		return true, nil
	default:
	}
	return false, nil
}
