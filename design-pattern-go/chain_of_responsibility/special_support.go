package chain_of_responsibility

type SpecialSupport struct {
	*Support
	number int
}

func NewSpecialSupport(name string, number int) *SpecialSupport {
	specialSupport := &SpecialSupport{
		Support: NewSupport(name),
		number:  number,
	}
	specialSupport.concreteSupport = specialSupport

	return specialSupport
}

func (s *SpecialSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumger() == s.number {
		return true
	} else {
		return false
	}
}
