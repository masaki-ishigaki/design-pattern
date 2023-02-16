package observer

import (
	"fmt"
	"time"
)

type Observer interface {
	Update(generator INumberGenerator)
}

type DigitObserver struct {
	concreteObserver Observer
}

func NewDigitObserver() *DigitObserver {
	digitObserver := &DigitObserver{}
	digitObserver.concreteObserver = digitObserver

	return digitObserver
}

func (d *DigitObserver) Update(generator INumberGenerator) {
	fmt.Printf("DigitObserver:%d\n", generator.GetNumber())
	time.Sleep(time.Microsecond * 100)
}

type GraphObserver struct {
	concreteObserver Observer
}

func NewGraphObserver() *GraphObserver {
	graphObserver := &GraphObserver{}
	graphObserver.concreteObserver = graphObserver

	return graphObserver
}
