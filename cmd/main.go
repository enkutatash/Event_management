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
	authRepo := repository.NewAuthRepo(dbConn.Db)
	authUsecase := usecase.NewAuthUseCase(authRepo)
	authhandler := handler.NewHandler(authUsecase)
	
	userRepo := repository.NewUserRepo(dbConn.Db,dbConn.Cache)
	userUsecase := usecase.NewuserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)


	route.InitRoute(r,authhandler,userHandler)


	adminRepo := repository.NewAdminRepo(dbConn.Db)
	adminUsecase := usecase.AdminUsecase(adminRepo)
	adminHandler := handler.NewAdminHandler(adminUsecase)
	route.InitAdminRoute(r,adminHandler)
	route.Run(r,":8081")

}