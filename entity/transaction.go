package entity

import (
	"time"

	"github.com/google/uuid"
)

//Transaction data
type Transaction struct {
	ID        ID
	Title     string
	Value     int64
	User      ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewTransaction create a new transaction
func NewTransaction(title string, value int64, user ID) (*Transaction, error) {
	b := &Transaction{
		ID:        NewID(),
		Title:     title,
		Value:     value,
		User:      user,
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
	if b.Title == "" || b.Value == 0 || b.User == uuid.Nil {
		return ErrInvalidEntity
	}
	return nil
}
