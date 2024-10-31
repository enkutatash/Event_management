package route

import (
	"event/handler"
	"event/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine,authhandler handler.AuthRequest,userhandler handler.UserHandler){
	// r.POST("/validate",userhandler.Validate)
	r.GET("/verify",authhandler.VerifyEmail)
	r.POST("/signup",authhandler.Signup)
	r.POST("/login",authhandler.Login)
	r.GET("/logout",authhandler.Logout)
	r.GET("/events",userhandler.GetAllEvents)
	r.GET("/event/:event_id",userhandler.GetEventById)
	r.GET("/book",userhandler.BookTicket)
}

func InitAdminRoute(r *gin.Engine,adminHandler handler.AdminHandler){
	
	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.AuthenticateAdmin())
	{
		adminRoute.POST("/create",adminHandler.CreateEvent)
		adminRoute.DELETE("/cancel/:event_id",adminHandler.CancelEvent)
	}
}


func Run(r *gin.Engine,address string) error{
	return r.Run(address)
}