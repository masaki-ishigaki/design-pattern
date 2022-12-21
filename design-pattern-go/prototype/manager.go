package prototype

type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{
		showcase: map[string]Product{},
	}
}

func (m *Manager) Register(name string, prototype Product) {
	m.showcase[name] = prototype
}

func (m *Manager) Create(prototypeName string) Product {
	p := m.showcase[prototypeName]
	return p.CreateCopy()
}
