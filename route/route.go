package route

import (
	"event/handler"
	"event/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine,userhandler handler.AuthRequest){
	
	r.POST("/signup",userhandler.Signup)
	r.POST("/login",userhandler.Login)
	r.GET("/logout",userhandler.Logout)
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