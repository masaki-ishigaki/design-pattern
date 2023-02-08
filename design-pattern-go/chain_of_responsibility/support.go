package chain_of_responsibility

import "fmt"

type ISupport interface {
	Resolve(trouble *Trouble) bool
}

type Support struct {
	name            string
	next            *Support
	concreteSupport ISupport
}

func NewSupport(name string) *Support {
	return &Support{
		name: name,
	}
}

func (s *Support) SetNext(next *Support) *Support {
	s.next = next
	return next
}

func (s *Support) Support(trouble *Trouble) {
	if s.concreteSupport.Resolve(trouble) {
		s.Done(trouble)
	} else if s.next != nil {
		s.next.Support(trouble)
	} else {
		s.Fail(trouble)
	}
}

func (s *Support) Done(trouble *Trouble) {
	fmt.Printf("%v is resolved by %v.\n", trouble, s)
}

func (s *Support) Fail(trouble *Trouble) {
	fmt.Printf("%v cannot be resolved.\n", trouble)
}

func (s *Support) String() string {
	return fmt.Sprintf("[%s]", s.name)
}
