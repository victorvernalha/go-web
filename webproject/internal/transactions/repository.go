package transactions

import "errors"

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(Transaction) error
}

type InMemoryRepository struct {
	Ts map[int]Transaction
}

func (r *InMemoryRepository) GetAll() ([]Transaction, error) {
	ts := []Transaction{}

	for _, t := range r.Ts {
		ts = append(ts, t)
	}

	return ts, nil
}

func (r *InMemoryRepository) Store(t Transaction) error {
	if _, exists := r.Ts[t.ID]; exists {
		return errors.New("transaction already exists")
	}
	r.Ts[t.ID] = t
	return nil
}
