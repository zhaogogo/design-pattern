package pkg

import (
	"errors"
	"fmt"
)

type One struct {
	Next
}

func (o *One) Do(s *slience) error {
	fmt.Println("one do")
	return nil
}

type Two struct {
	Next
}

func (o *Two) Do(s *slience) error {
	fmt.Println("Two do")
	return nil
}

type Thread struct {
	Next
}

func (o *Thread) Do(s *slience) error {
	fmt.Println("Thread do")
	return errors.New("Thread error")
}

type Four struct {
	Next
}

func (o *Four) Do(s *slience) error {
	fmt.Println("Four do")
	return nil
}

type Five struct {
	Next
}

func (o *Five) Do(s *slience) error {
	fmt.Println("Five do")
	return nil
}
