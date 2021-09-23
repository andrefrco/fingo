package transaction

import (
	"github.com/andrefrco/gofin/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Transaction, error)
	Search(query string) ([]*entity.Transaction, error)
	List() ([]*entity.Transaction, error)
}

//Writer Transaction writer
type Writer interface {
	Create(e *entity.Transaction) (entity.ID, error)
	Update(e *entity.Transaction) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetTransaction(id entity.ID) (*entity.Transaction, error)
	SearchTransactions(query string) ([]*entity.Transaction, error)
	ListTransactions() ([]*entity.Transaction, error)
	CreateTransaction(title string, value int64) (entity.ID, error)
	UpdateTransaction(e *entity.Transaction) error
	DeleteTransaction(id entity.ID) error
}
