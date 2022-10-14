package main

import (
	"github.com/fajartd02/golang-devcamp/config"
	"github.com/fajartd02/golang-devcamp/core/entity"
	"github.com/fajartd02/golang-devcamp/core/module"
	"github.com/fajartd02/golang-devcamp/handler"
	repository "github.com/fajartd02/golang-devcamp/repository"
	"github.com/fajartd02/golang-devcamp/routes"
)

func main() {

	db := config.Init()
	db.AutoMigrate(&entity.Product{}, &entity.Variant{}, &entity.User{}, &entity.Comment{})

	productRepo := repository.New()
	productUc := module.NewProductUsecase(productRepo)
	productHdl := handler.NewProductHandler(productUc)

	userRepo := repository.NewUser()
	userUc := module.NewUserUsecase(userRepo)
	userHdl := handler.NewUserHandler(userUc)

	r := routes.SetupRoutes(db, *productHdl, *userHdl)
	r.Run(":8080")
}
