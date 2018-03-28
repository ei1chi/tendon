package tendon

import (
	"errors"

	et "github.com/hajimehoshi/ebiten"
)

type Scene interface {
	Load()
	Update(*et.Image) (Scene, error)
	ParentScene() Scene
}

func RootScene(s Scene) Scene {
	for s.ParentScene() != nil {
		s = s.ParentScene()
	}
	return s
}

type SceneBase struct {
	next, Parent Scene
	task         chan struct{}
	hasLoaded    bool
}

func (s *SceneBase) ParentScene() Scene {
	return s.Parent
}

func (s *SceneBase) StartNextLoading(next Scene) error {
	if s.task != nil {
		return errors.New("next loading has already started!")
	}
	if s.next == nil {
		return errors.New("next has not created!")
	}
	s.task = make(chan struct{}, 1)
	s.next = next
	go func() {
		s.next.Load()
		<-s.task
	}()
	return nil
}

func (s *SceneBase) NextScene() (Scene, error) {
	if s.task == nil {
		return nil, errors.New("hasn't started loading")
	}
	if s.next == nil {
		return nil, errors.New("next has not created!")
	}
	if s.hasLoaded {
		return s.next, nil
	}
	select {
	case s.task <- struct{}{}: // 空である
		s.hasLoaded = true
		return s.next, nil
	default:
	}
	return nil, nil
}
