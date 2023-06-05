package transactions

import (
	"errors"
	"time"
)

type Service interface {
	GetAll() ([]Transaction, error)
	Create(code, currency, sender, receiver string, amount float64, date time.Time) (Transaction, error)
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

func (s *DefaultService) generateNewID() (int, error) {
	if ts, err := s.Repo.GetAll(); err == nil {
		return len(ts), nil
	}
	return 0, errors.New("could not generate new ID")
}
