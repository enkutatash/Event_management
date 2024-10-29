package util

import (
	"event/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Id int
	FullName string
	Email    string
	Role     string
	jwt.StandardClaims
}

var Secret_key = os.Getenv("SECRET_KEY")


func GenerateToken(id int,name string,email string,role string) (*models.LoginRes, error){

	claim := &UserClaim{
		Id: id,
		FullName: name,
		Email: email,
		Role:role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().AddDate(1, 0, 0).Unix(),
		},

	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(Secret_key))
	if err != nil {
		return nil, err
	}
	user := &models.LoginRes{
		Id:		 id,
		AccessToken: token,
		FullName:   name,
		Email: email,
	}
	return user, nil
}

func VerifyToken(tokenString string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret_key), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaim)
	if !ok {
		return nil, err
	}
	return claims, nil
}
