package transaction

import (
	"time"

	"github.com/andrefrco/gofin/entity"
)

//Service transaction usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateTransaction create a transaction
func (s *Service) CreateTransaction(title string, value int64) (entity.ID, error) {
	b, err := entity.NewTransaction(title, value)
	if err != nil {
		return b.ID, err
	}
	return s.repo.Create(b)
}

//GetTransaction get a transaction
func (s *Service) GetTransaction(id entity.ID) (*entity.Transaction, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return b, nil
}

//SearchTransactions search transactions
func (s *Service) SearchTransactions(query string) ([]*entity.Transaction, error) {
	transactions, err := s.repo.Search(query)
	if err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, entity.ErrNotFound
	}
	return transactions, nil
}

//ListTransactions list transactions
func (s *Service) ListTransactions() ([]*entity.Transaction, error) {
	transactions, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, entity.ErrNotFound
	}
	return transactions, nil
}

//DeleteTransaction Delete a transaction
func (s *Service) DeleteTransaction(id entity.ID) error {
	_, err := s.GetTransaction(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateTransaction Update a transaction
func (s *Service) UpdateTransaction(e *entity.Transaction) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
