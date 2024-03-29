package chain_of_responsibility

import "fmt"

type Trouble struct {
	number int
}

func NewTrouble(number int) *Trouble {
	return &Trouble{
		number,
	}
}

func (t *Trouble) GetNumger() int {
	return t.number
}

func (t *Trouble) String() string {
	return fmt.Sprintf("[Trouble %d]", t.number)
}
