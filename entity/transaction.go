package entity

import (
	"time"
)

//Transaction data
type Transaction struct {
	ID        ID
	Title     string
	Value     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewTransaction create a new transaction
func NewTransaction(title string, value int64) (*Transaction, error) {
	b := &Transaction{
		ID:        NewID(),
		Title:     title,
		Value:     value,
		CreatedAt: time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

//Validate validate transaction
func (b *Transaction) Validate() error {
	if b.Title == "" || b.Value == 0 {
		return ErrInvalidEntity
	}
	return nil
}
