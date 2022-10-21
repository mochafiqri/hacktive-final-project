package main

import (
	"finalProject/handler"
	"finalProject/infra"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := infra.RouterInit()
	db := infra.DbInit()

	var h = handler.NewHandler(db)

	handler.NewUserHandler(r, h)
	handler.NewPhotoHandler(r, h)
	handler.NewCommentHandler(r, h)
	handler.NewSocialMediaHandler(r, h)

	appPort := os.Getenv("PORT")

	if appPort == "" {
		log.Fatal("app port didn't define")
	}

	log.Fatal(r.Run(":" + appPort))

}
