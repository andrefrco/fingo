package presenter

import "github.com/andrefrco/gofin/entity"

// Transaction data
type Transaction struct {
	ID    entity.ID `json:"id"`
	Title string    `json:"title"`
	Value int64     `json:"value"`
}
