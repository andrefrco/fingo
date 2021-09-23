package transaction

import (
	"strings"

	"github.com/andrefrco/gofin/entity"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Transaction
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Transaction{}
	return &inmem{
		m: m,
	}
}

//Create a transaction
func (r *inmem) Create(e *entity.Transaction) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a transaction
func (r *inmem) Get(id entity.ID) (*entity.Transaction, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a transaction
func (r *inmem) Update(e *entity.Transaction) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search transactions
func (r *inmem) Search(query string) ([]*entity.Transaction, error) {
	var d []*entity.Transaction
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Title), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List transactions
func (r *inmem) List() ([]*entity.Transaction, error) {
	var d []*entity.Transaction
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a transaction
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
