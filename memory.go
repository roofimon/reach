package reach

type Memory struct{}

func (m *Memory) findAll() ([]Info, error) {
	return []Info{
		Info{Name: "roof", Email: "roofimon@gmail.com"},
		Info{Name: "foor", Email: "foorimon@gmail.com"},
	}, nil
}
