package main

import (
	routes "file_uploading/Routes"
	"file_uploading/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	repository.DbCon()
	r := gin.Default()
	routes.AllRoutes(r)

	log.Println("stating server .......")
	r.Run()

}
