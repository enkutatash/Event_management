package handler

import (
	"event/models"
	"event/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRes struct{
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type AuthRequest interface {
	Login(c *gin.Context) 
	Signup(c *gin.Context) 
	Logout(c *gin.Context)
	VerifyEmail(c *gin.Context) 
}


type AuthHandler struct {
	AuthUsecase usecase.AuthUsecase
}

// VerifyEmail implements UserHandler.
func (u AuthHandler) VerifyEmail(c *gin.Context) {
	email := c.Query("email")
	token := c.Query("token")
	
	err := u.AuthUsecase.VerifyEmail(&email, &token)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"message": email + " verified" })
}

// Login implements AuthRequest.
func (a AuthHandler) Login(c *gin.Context)  {
	var user models.UserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	u,err := a.AuthUsecase.Login(user.Email,user.Password)
	idString := strconv.Itoa(u.Id)
	c.SetCookie("user_id",idString,3600,"","",false,true)
	
	
	res := &UserRes{
		FullName: u.FullName,
		Email: u.Email,
	}

	if err != nil{
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Header("Authorization", "Bearer " + u.AccessToken)
	c.IndentedJSON(200, gin.H{"user": res})
}

// Logout implements AuthRequest.
func (a AuthHandler) Logout(c *gin.Context)  {
	c.SetCookie("jwt","",-1,"","",false,true)
	c.JSON(200,gin.H{"message":"logout success"})
}

// Signup implements AuthRequest.
func (a AuthHandler) Signup(c *gin.Context)  {
	var user models.UserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	id,err := a.AuthUsecase.Signup(user.FullName,user.Email,user.Password)
	if err != nil{
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"message": "success","id":id})
}



func NewHandler(authusecase usecase.AuthUsecase) AuthRequest {
	return AuthHandler{
		AuthUsecase: authusecase,
	}
}
