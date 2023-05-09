package hunt

import (
	"errors"
)

// Errors
var (
	ErrTired = errors.New("cannot hunt, i am really tired")
	ErrNotHungry = errors.New("cannot hunt, i am not hungry")
	ErrPreyEscaped = errors.New("could not catch it")
	ErrNotPrey = errors.New("prey cannot be nil")

)

type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

type Prey struct {
	name  string
	speed int
}

func (s *Shark) Hunt(p *Prey) error {
	if p == nil {
		return ErrNotPrey
	}
	if s.tired {
		return ErrTired
	}
	if !s.hungry {
		return ErrNotHungry
	}
	if p.speed >= s.speed {
		s.tired = true
		return ErrPreyEscaped
	} 

	s.hungry = false
	s.tired = true
	return nil
}
