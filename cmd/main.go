package main

import (
	"event/db"
	"event/handler"
	"event/repository"
	"event/route"
	"event/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("couldn't init db", err)
	}
	
		log.Println("DB connected")

	r := gin.New()
	r.Use(gin.Logger())
	userRepo := repository.NewAuthRepo(dbConn.Db)
	userUsecase := usecase.NewUseCase(userRepo)
	userhandler := handler.NewHandler(userUsecase)
	route.InitRoute(r,userhandler)


	adminRepo := repository.NewAdminRepo(dbConn.Db)
	adminUsecase := usecase.AdminUsecase(adminRepo)
	adminHandler := handler.NewAdminHandler(adminUsecase)
	route.InitAdminRoute(r,adminHandler)
	route.Run(r,":8080")
}