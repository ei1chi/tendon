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
	slave Scene
	task  chan struct{}
}

func (s *SceneBase) StartSlaveLoading(slave Scene) error {
	if s.task != nil {
		return errors.New("slave Scene has already created!!!")
	}
	s.slave = slave
	s.task = make(chan struct{}, 1)
	go func() {
		s.slave.Load()
		<-s.task
	}()
	return nil
}

func (s *SceneBase) HasSlaveLoaded() (bool, error) {
	if s.task == nil {
		return false, errors.New("hasn't started loading")
	}
	select {
	case s.task <- struct{}{}: // 空である
		return true, nil
	default:
	}
	return false, nil
}
