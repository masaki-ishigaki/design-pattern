package observer

import (
	"math/rand"
	"time"
)

type INumberGenerator interface {
	GetNumber() int
	Execute()
}

type NumberGenerator struct {
	observers         []Observer
	concreteGenerator INumberGenerator
}

func (n *NumberGenerator) AddObserver(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NumberGenerator) NotifyObservers() {
	for _, o := range n.observers {
		o.Update(n.concreteGenerator)
	}
}

type RandomNumberGenerator struct {
	number          int
	numberGenerator NumberGenerator
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	randomNumberGenerator := &RandomNumberGenerator{
		number: 0,
	}
	randomNumberGenerator.numberGenerator.concreteGenerator = randomNumberGenerator

	return randomNumberGenerator
}

func (r *RandomNumberGenerator) GetNumber() int {
	return r.number
}

func (r *RandomNumberGenerator) Execute() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		r.number = rand.Intn(50)
		r.numberGenerator.NotifyObservers()
	}
}
