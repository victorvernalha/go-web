package transactions

import "errors"

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(Transaction) error
}

type InMemoryRepository struct {
	Ts map[int]Transaction
}

func CreateInMemoryRepo() InMemoryRepository {
	return InMemoryRepository{make(map[int]Transaction)}
}

func (r *InMemoryRepository) exists(t *Transaction) (exists bool) {
	_, exists = r.Ts[t.ID]
	return
}

func (r *InMemoryRepository) GetAll() ([]Transaction, error) {
	ts := []Transaction{}

	for _, t := range r.Ts {
		ts = append(ts, t)
	}

	return ts, nil
}

func (r *InMemoryRepository) Store(t Transaction) error {
	if r.exists(&t) {
		return errors.New("transaction already exists")
	}
	r.Ts[t.ID] = t
	return nil
}
