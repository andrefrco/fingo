package entity_test

import (
	"testing"

	"github.com/andrefrco/gofin/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	b, err := entity.NewTransaction("Supermarket", 100)
	assert.Nil(t, err)
	assert.Equal(t, b.Title, "Supermarket")
	assert.NotNil(t, b.ID)
}

func TestTransactionValidate(t *testing.T) {
	type test struct {
		title string
		value int64
		want  error
	}

	tests := []test{
		{
			title: "Supermarket",
			value: 100,
			want:  nil,
		},
		{
			title: "Supermarket",
			value: 0,
			want:  entity.ErrInvalidEntity,
		},
		{
			title: "",
			value: 100,
			want:  entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewTransaction(tc.title, tc.value)
		assert.Equal(t, err, tc.want)
	}

}
