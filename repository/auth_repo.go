package repository

import (
	"context"
	"database/sql"
	"errors"
	"event/models"

)


type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string,args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type AuthRepository interface {
	GetUserByEmail(email string) (*models.UserRes, error)
	CreateUser(user *models.UserReq) (*int,error)
}

type authRepo struct {
	db *sql.DB
}

// CreateUser implements AuthRepository.
func (a *authRepo) CreateUser(user *models.UserReq) (*int,error) {
	c,cancel := context.WithCancel(context.Background())
	defer cancel()
	var userId int
	query := "INSERT INTO users(fullname,email,password) VALUES($1,$2,$3)  RETURNING id"
	// _, err := a.db.ExecContext(c, query, user.FullName, user.Email, user.Password)
	err :=a.db.QueryRowContext(c,query,user.FullName,user.Email,user.Password).Scan(&userId)
	// err := a.db.QueryRowContext(c,query,user.FullName,user.Email,user.Password)
	if err != nil{
		return nil,errors.New("failed to create user")
	}
	
	return &userId,nil
}

// GetUserByEmail implements AuthRepository.
func (a *authRepo) GetUserByEmail(email string) (*models.UserRes, error) {
	c,cancel := context.WithCancel(context.Background())
	defer cancel()
	var user models.UserRes;
	query := "SELECT id,fullname, email,password FROM users WHERE email = $1"
	err := a.db.QueryRowContext(c,query,email).Scan(&user.Id,&user.FullName,&user.Email,&user.Password)
	if err != nil{
		return nil,nil
	}
	return &user,nil
}

func NewAuthRepo(db *sql.DB) AuthRepository {
	return &authRepo{
		db: db,
	}
}
