package reach

type Provider interface {
	findAll() ([]Info, error)
}
