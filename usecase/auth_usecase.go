package usecase

import (
	"errors"
	"event/models"
	"event/repository"
	"event/util"
	
)

const (
	SECRETKEY = "secret"
)

type AuthUsecase interface {
	Signup(fullname string, email string, password string) (*int,error)
	Login(email string, password string) (*models.LoginRes, error)
}

type authUseCase struct {
	AuthRepo repository.AuthRepository
}

// Login implements AuthUsecase.
func (a *authUseCase) Login(email string, password string) (*models.LoginRes, error) {
	user,err := a.AuthRepo.GetUserByEmail(email)
	if err != nil{
		return nil,err
	}

	err = util.CheckPassword(user.Password,password)
	if err != nil{
		return nil,errors.New("password is incorrect")
	}

	var role string
	if user.Email == "admin@com"{
		role = "admin"
	}else{
		role = "user"
	}
	
	res , err := util.GenerateToken(user.Id,user.FullName,user.Email,role)

	if err != nil{
		return nil,err
	}

	return res, nil

}

// Signup implements AuthUsecase.
func (a *authUseCase) Signup(fullname string, email string, password string) (*int,error) {

	// check if email already exist 
	us,err := a.AuthRepo.GetUserByEmail(email)

	if err != nil{
		return nil,err
	}
	if us != nil{
		return nil,errors.New("email already exist")
	}
	
	// hash password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil,errors.New("failed to hash password")
	}
	
	user := &models.UserReq{
		FullName: fullname,
		Email:    email,
		Password: hashedPassword,
	}

	id,err := a.AuthRepo.CreateUser(user)
	if err != nil{
		return nil,err
	}
	return id,nil
	// call repository
}

func NewUseCase(authRepo repository.AuthRepository) AuthUsecase {
	return &authUseCase{
		AuthRepo: authRepo,
	}
}
