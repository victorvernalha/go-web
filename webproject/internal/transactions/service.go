package transactions

import (
	"errors"
	"time"
)

type Service interface {
	GetAll() ([]Transaction, error)
	Create(code, currency, sender, receiver string, amount float64, date time.Time) (Transaction, error)
	Replace(t Transaction) (Transaction, error)
}

type DefaultService struct {
	Repo Repository
}

func (s *DefaultService) GetAll() ([]Transaction, error) {
	return s.Repo.GetAll()
}

func (s *DefaultService) Create(code, currency, sender, receiver string, amount float64, date time.Time) (t Transaction, err error) {
	id, err := s.generateNewID()
	if err != nil {
		return
	}

	t = Transaction{id, code, currency, amount, sender, receiver, date}
	err = s.Repo.Store(t)
	if err != nil {
		err = errors.New("error saving new transaction")
		return
	}
	return
}

func (s *DefaultService) Replace(t Transaction) (Transaction, error) {
	if err := s.Repo.Replace(t); err != nil {
		return t, err
	}
	return t, nil
}

func (s *DefaultService) generateNewID() (int, error) {
	if ts, err := s.Repo.GetAll(); err == nil {
		newId := 0
		for _, t := range ts {
			newId = Max(t.ID, newId) + 1
		}
		return newId, nil
	}
	return 0, errors.New("could not generate new ID")
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
