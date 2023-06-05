package transactions

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(Transaction) error
}

type InMemoryRepository struct {
	Ts []Transaction
}

func (r *InMemoryRepository) GetAll() ([]Transaction, error) {
	return r.Ts, nil
}

func (r *InMemoryRepository) Store(t Transaction) error {
	r.Ts = append(r.Ts, t)
	return nil
}
