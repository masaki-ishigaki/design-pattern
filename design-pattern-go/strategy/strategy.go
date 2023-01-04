package strategy

import (
	"math/rand"
	"time"
)

type Strategy interface {
	NextHand() *Hand
	Study(win bool)
}

/*****************************
 * WinningStrategy
 ****************************/
type WinningStrategy struct {
	won      bool
	prevHand Hand
}

func NewWinningStrategy() *WinningStrategy {
	return &WinningStrategy{
		won:      false,
		prevHand: *NewHand(HAND_KIND_ROCK),
	}
}

func (w *WinningStrategy) NextHand() *Hand {
	rand.Seed(time.Now().UnixNano())
	if !w.won {
		w.prevHand = *NewHand(GetHandKindById(rand.Intn(2)))
	}
	return &w.prevHand
}

func (w *WinningStrategy) Study(win bool) {
	w.won = win
}

/*****************************
 * ProbStrategy
 ****************************/
type ProbStrategy struct {
	prevHandValue    int
	currentHandValue int
	history          [][]int
}

func NewProbStrategy() *ProbStrategy {
	return &ProbStrategy{
		prevHandValue:    HAND_KIND_ROCK.id,
		currentHandValue: HAND_KIND_ROCK.id,
		history: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
}

func (p *ProbStrategy) NextHand() *Hand {
	rand.Seed(time.Now().UnixNano())
	bet := rand.Intn(p.getSum(p.currentHandValue))
	var handValue int
	if bet < p.history[p.currentHandValue][0] {
		handValue = HAND_KIND_ROCK.id
	} else if bet < (p.history[p.currentHandValue][0] + p.history[p.currentHandValue][1]) {
		handValue = HAND_KIND_SCISSORS.id
	} else {
		handValue = HAND_KIND_PAPER.id
	}

	p.prevHandValue = p.currentHandValue
	p.currentHandValue = handValue

	return NewHand(GetHandKindById(handValue))
}

func (p *ProbStrategy) Study(win bool) {
	if win {
		p.history[p.prevHandValue][p.currentHandValue] += 1
	} else {
		p.history[p.prevHandValue][(p.currentHandValue+1)%3] += 1
		p.history[p.prevHandValue][(p.currentHandValue+2)%3] += 1
	}
}

func (p *ProbStrategy) getSum(id int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += p.history[id][i]
	}
	return sum
}
