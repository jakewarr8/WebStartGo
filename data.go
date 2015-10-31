package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewOpen(dt string, c string) (DB, error) {
	db, err := sql.Open(dt, c)
	return DB{db}, err
}
