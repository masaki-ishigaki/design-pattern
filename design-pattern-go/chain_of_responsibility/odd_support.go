package chain_of_responsibility

type OddSupport struct {
	*Support
}

func NewOddSupport(name string) *OddSupport {
	oddSupport := &OddSupport{
		Support: NewSupport(name),
	}
	oddSupport.concreteSupport = oddSupport

	return oddSupport
}

func (o *OddSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumger()%2 == 1 {
		return true
	} else {
		return false
	}
}
