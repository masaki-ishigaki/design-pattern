package chain_of_responsibility

type LimitSupport struct {
	*Support
	limit int
}

func NewLimitSupport(name string, limit int) *LimitSupport {
	limitSupport := &LimitSupport{
		Support: NewSupport(name),
		limit:   limit,
	}
	limitSupport.concreteSupport = limitSupport

	return limitSupport
}

func (l *LimitSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumger() < l.limit {
		return true
	} else {
		return false
	}
}
