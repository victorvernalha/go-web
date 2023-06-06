package transactions

import (
	"errors"
	"fmt"
)

type InMemoryRepository struct {
	Ts map[int]Transaction
}

func CreateInMemoryRepo() InMemoryRepository {
	return InMemoryRepository{make(map[int]Transaction)}
}

func (r *InMemoryRepository) exists(t *Transaction) (exists bool) {
	return r.existsID(t.ID)
}

func (r *InMemoryRepository) existsID(id int) (exists bool) {
	_, exists = r.Ts[id]
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

func (r *InMemoryRepository) Delete(id int) error {
	if !r.existsID(id) {
		return IDDoesNotExistError(id)
	}
	delete(r.Ts, id)
	return nil
}

func (r *InMemoryRepository) Replace(t Transaction) error {
	if !r.exists(&t) {
		return IDDoesNotExistError(t.ID)
	}
	r.Ts[t.ID] = t
	return nil
}

func (r *InMemoryRepository) Find(id int) (Transaction, error) {
	if !r.existsID(id) {
		return Transaction{}, IDDoesNotExistError(id)
	}
	return r.Ts[id], nil
}

func IDDoesNotExistError(id int) error {
	return fmt.Errorf("transaction with ID %d does not exist", id)
}
