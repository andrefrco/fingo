package presenter

import (
	"github.com/andrefrco/gofin/entity"
)

//User data
type User struct {
	ID        entity.ID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
