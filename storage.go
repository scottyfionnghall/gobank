package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccout(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore(connStr string) (*PostgressStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{db: db}, nil
}

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgressStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account(
		id serial PRIMARY KEY,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		number integer,
		balance integer,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgressStore) CreateAccount(account *Account) error {
	stmt, err := s.db.Prepare(`
	INSERT INTO account(first_name,last_name,number,balance,created_at) 
	VALUES($1,$2,$3,$4,$5)
	`)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt,
	).Err()

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgressStore) DeleteAccout(int) error {
	return nil
}

func (s *PostgressStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgressStore) GetAccountById(int) (*Account, error) {
	return nil, nil
}
