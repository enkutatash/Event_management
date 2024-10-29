package handler

import (
	"event/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetAllEvents(c *gin.Context)
	GetEventById(c *gin.Context)
	BookTicket(c *gin.Context)
}

type userHandler struct {
	UserUsecase usecase.UserUsecase
}

// BookTicket implements UserHandler.
func (u *userHandler) BookTicket(c *gin.Context) {
	panic("unimplemented")
}

// GetAllEvents implements UserHandler.
func (u *userHandler) GetAllEvents(c *gin.Context) {
	events, err := u.UserUsecase.GetAllEvents()
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"events": events})
}

// GetEventById implements UserHandler.
func (u *userHandler) GetEventById(c *gin.Context) {
	eventId := c.Param("event_id")
	event, err := u.UserUsecase.GetEventById(&eventId)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"event": event})
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{
		UserUsecase: uu,
	}
}
