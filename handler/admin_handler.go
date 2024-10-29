package handler

import (
	"event/models"
	"event/usecase"

	"github.com/gin-gonic/gin"
)

type AdminHandler interface {
	CreateEvent(c *gin.Context)
	CancelEvent(c *gin.Context)
}

type adminHandler struct {
	AdminUsecase usecase.AdminUsecase
}

// CancelEvent implements AdminHandler.
func (a *adminHandler) CancelEvent(c *gin.Context) {
	eventId := c.Param("event_id")
	err := a.AdminUsecase.CancelEvent(&eventId)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"message": "event cancelled"})
}

// CreateEvent implements AdminHandler.
func (a *adminHandler) CreateEvent(c *gin.Context) {
	var newEvent models.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	id,err := a.AdminUsecase.CreateEvent(&newEvent)
	if err!= nil{
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"event_id": id})
}

func NewAdminHandler(au usecase.AdminUsecase) AdminHandler {
	return &adminHandler{
		AdminUsecase: au,
	}
}
