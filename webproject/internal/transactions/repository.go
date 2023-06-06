package transactions

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(Transaction) error
	Replace(Transaction) error
	Find(int) (Transaction, error)
	Delete(int) error
}
