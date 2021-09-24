package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andrefrco/gofin/entity"
	"github.com/jmoiron/sqlx"
)

//User repo
type User struct {
	db *sqlx.DB
}

//NewUser create new repository
func NewUser(db *sql.DB) *User {
	var dbx = sqlx.NewDb(db, "postgres")
	return &User{
		db: dbx,
	}
}

//Create an user
func (r *User) Create(e *entity.User) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into user (id, email, password, first_name, last_name, created_at) 
		values($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Email,
		e.Password,
		e.FirstName,
		e.LastName,
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

//Get an user
func (r *User) Get(id entity.ID) (*entity.User, error) {
	return getUser(id, r.db.DB)
}

func getUser(id entity.ID, db *sql.DB) (*entity.User, error) {
	stmt, err := db.Prepare(`select id, email, first_name, last_name, created_at from user where id = $1`)
	if err != nil {
		return nil, err
	}
	var u entity.User
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &u, nil
}

//Update an user
func (r *User) Update(e *entity.User) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update user set email = $1, password = $2, first_name = $3, last_name = $4, updated_at = $5 where id = $6", e.Email, e.Password, e.FirstName, e.LastName, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search users
func (r *User) Search(query string) ([]*entity.User, error) {
	stmt, err := r.db.Prepare(`select id from user where name like '%' || $1 || '%'`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []entity.ID
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i entity.ID
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var users []*entity.User
	for _, id := range ids {
		u, err := getUser(id, r.db.DB)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

//List users
func (r *User) List() ([]*entity.User, error) {
	stmt, err := r.db.Prepare(`select id from user`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []entity.ID
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i entity.ID
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var users []*entity.User
	for _, id := range ids {
		u, err := getUser(id, r.db.DB)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

//Delete an user
func (r *User) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from user where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
