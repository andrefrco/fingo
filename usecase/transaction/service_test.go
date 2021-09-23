package transaction

import (
	"testing"
	"time"

	"github.com/andrefrco/gofin/entity"

	"github.com/stretchr/testify/assert"
)

func newFixtureTransaction() *entity.Transaction {
	return &entity.Transaction{
		Title:     "Drink's distributor",
		Value:     100,
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureTransaction()
	_, err := m.CreateTransaction(u.Title, u.Value)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureTransaction()
	u2 := newFixtureTransaction()
	u2.Title = "Supermarket"

	uID, _ := m.CreateTransaction(u1.Title, u1.Value)
	_, _ = m.CreateTransaction(u2.Title, u2.Value)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchTransactions("distributor")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Drink's distributor", c[0].Title)

		c, err = m.SearchTransactions("iFood")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListTransactions()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetTransaction(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Title, saved.Title)
	})
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureTransaction()
	id, err := m.CreateTransaction(u.Title, u.Value)
	assert.Nil(t, err)
	saved, _ := m.GetTransaction(id)
	saved.Title = "Supermarket"
	assert.Nil(t, m.UpdateTransaction(saved))
	updated, err := m.GetTransaction(id)
	assert.Nil(t, err)
	assert.Equal(t, "Supermarket", updated.Title)
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureTransaction()
	u2 := newFixtureTransaction()
	u2ID, _ := m.CreateTransaction(u2.Title, u2.Value)

	err := m.DeleteTransaction(u1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteTransaction(u2ID)
	assert.Nil(t, err)
	_, err = m.GetTransaction(u2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
