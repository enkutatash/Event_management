package handler

import (
	"event/usecase"
	"fmt"
	"strconv"

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
	event_id := c.Query("event_id")
	ticket_no := c.Query("ticket_no")
	user_id, err := c.Cookie("user_id")
	fmt.Println("cookie", user_id)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "user not logged in"})
		return
	}

	ticketNoInt, err := strconv.Atoi(ticket_no)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid ticket number"})
		return
	}
	userIDInt, err := strconv.Atoi(user_id)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid user ID"})
		return
	}
	err = u.UserUsecase.BookTicket(&event_id, &userIDInt, &ticketNoInt)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"message": "ticket booked"})

}

type Pagination struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

// GetAllEvents implements UserHandler.
func (u *userHandler) GetAllEvents(c *gin.Context) {
	var p Pagination
	if err := c.ShouldBindQuery(&p); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if p.Offset == 0 {
		p.Offset = 1
	}
	events, err := u.UserUsecase.GetAllEvents(p.Offset, p.Limit)
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
