package tendon

type Stm struct {
	state, count int
}

func (s Stm) Update() {
	count += 1
}

func (s Stm) Elapsed() int {
	return s.count
}

func (s Stm) Transition(next int) {
	s.state = next
	s.count = 0
}

func (s Stm) Continue(next int) {
	s.state = next
}
