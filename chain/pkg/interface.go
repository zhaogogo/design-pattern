package pkg

import (
	"reflect"
)

type ChanHandler interface {
	SetNext(ChanHandler) ChanHandler
	Execute(*slience) error
	Do(*slience) error
}

type Next struct {
	nextHandler ChanHandler
}

func (n *Next) SetNext(handler ChanHandler) ChanHandler {
	n.nextHandler = handler
	return handler
}

func (n *Next) Execute(s *slience) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(s); err != nil {
			return err
		}
		t := reflect.TypeOf(n.nextHandler)
		if t.Kind() != reflect.Ptr {
			panic("nextHandler must be ptr")
		}
		s.executechan = append(s.executechan, t.Elem().Name())
		return n.nextHandler.Execute(s)
	}
	return nil
}

type slience struct {
	name            string
	starthandler    ChanHandler
	registerchanlen int
	executechan     []string
}

func (s *slience) Execute(si *slience) error {
	return s.starthandler.Execute(si)
}

func (s *slience) Slice() []string {
	return s.executechan
}

type startHandler struct {
	Next
}

func (s startHandler) Do(c *slience) error {
	return nil
}

func NewChan(name string, actions ...ChanHandler) *slience {
	s := &slience{
		name:            name,
		registerchanlen: len(actions),
		executechan:     make([]string, 0, len(actions)),
	}
	var handler ChanHandler = &startHandler{}
	s.starthandler = handler

	for _, action := range actions {
		handler = handler.SetNext(action)
	}

	return s
}
