package chain_of_responsibility

type NoSupport struct {
	*Support
}

func NewNoSupport(name string) *NoSupport {
	noSupport := &NoSupport{
		Support: NewSupport(name),
	}
	noSupport.concreteSupport = noSupport

	return noSupport
}

func (n *NoSupport) Resolve(trouble *Trouble) bool {
	return false
}
