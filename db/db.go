package db

import (
	"database/sql"


	"github.com/go-redis/redis/v8"

	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
	Cache *redis.Client
}


func NewDatabase ()(*Database, error) {
	// var opts *redis.Options
	db, err := sql.Open("postgres", "postgresql://postgres:enku0811@localhost:5433/users?sslmode=disable")
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:            "localhost:6379",
		Password:        "",
		DB:              0,
		
	})
	

	return &Database{Db: db,Cache: rdb}, nil
}

// func (d *Database) close(){
// 	d.Db.Close()
// }

// func (d *Database) GetDB() *sql.DB {
// 	return d.Db
// }