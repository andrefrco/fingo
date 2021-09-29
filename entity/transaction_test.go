package entity_test

import (
	"testing"

	"github.com/andrefrco/gofin/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	user := uuid.New()
	b, err := entity.NewTransaction("Supermarket", 100, user)
	assert.Nil(t, err)
	assert.Equal(t, b.Title, "Supermarket")
	assert.NotNil(t, b.ID)
	assert.Equal(t, b.User, user)
}

func TestTransactionValidate(t *testing.T) {
	type test struct {
		title string
		value int64
		user  uuid.UUID
		want  error
	}

	tests := []test{
		{
			title: "Supermarket",
			value: 100,
			user:  uuid.New(),
			want:  nil,
		},
		{
			title: "Supermarket",
			value: 0,
			user:  uuid.New(),
			want:  entity.ErrInvalidEntity,
		},
		{
			title: "",
			value: 100,
			user:  uuid.New(),
			want:  entity.ErrInvalidEntity,
		},
		{
			title: "Supermarket",
			value: 100,
			user:  uuid.UUID{},
			want:  entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewTransaction(tc.title, tc.value, tc.user)
		assert.Equal(t, err, tc.want)
	}

}
