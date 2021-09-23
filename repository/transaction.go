package repository

import (
	"database/sql"
	"time"

	"github.com/andrefrco/gofin/entity"
	"github.com/jmoiron/sqlx"
)

//Transaction postgres repo
type Transaction struct {
	db *sqlx.DB
}

//NewTransaction create new repository
func NewTransaction(db *sql.DB) *Transaction {
	var dbx = sqlx.NewDb(db, "postgres")
	return &Transaction{
		db: dbx,
	}
}

//Create a transaction
func (r *Transaction) Create(e *entity.Transaction) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into transaction (id, title, value, created_at) 
		values($1,$2,$3,$4)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Title,
		e.Value,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

//Get a transaction
func (r *Transaction) Get(id entity.ID) (*entity.Transaction, error) {
	var b entity.Transaction
	err := r.db.QueryRow(
		`select id, title, value, created_at from transaction where id = $1`,
		id,
	).Scan(&b.ID, &b.Title, &b.Value, &b.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &b, nil
}

//Update a transaction
func (r *Transaction) Update(e *entity.Transaction) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update transaction set title = $1, value = $2, updated_at = $3 where id = $4", e.Title, e.Value, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search transactions
func (r *Transaction) Search(query string) ([]*entity.Transaction, error) {
	stmt, err := r.db.Prepare(`select id, title, value, created_at from transaction where title like '%' || $1 || '%'`)
	if err != nil {
		return nil, err
	}
	var transactions []*entity.Transaction
	rows, err := stmt.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Transaction
		err = rows.Scan(&b.ID, &b.Title, &b.Value, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &b)
	}
	return transactions, nil
}

//List transactions
func (r *Transaction) List() ([]*entity.Transaction, error) {
	stmt, err := r.db.Prepare(`select id, title, value, created_at from transaction`)
	if err != nil {
		return nil, err
	}
	var transactions []*entity.Transaction
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Transaction
		err = rows.Scan(&b.ID, &b.Title, &b.Value, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &b)
	}
	return transactions, nil
}

//Delete a transaction
func (r *Transaction) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from transaction where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
