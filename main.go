package main

import (
	"fmt"
	"github.com/MickStanciu/SC/src/controller"
	"github.com/MickStanciu/SC/src/repository"
	"github.com/MickStanciu/SC/src/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("HEY")

	jediRepository := repository.NewRealDB()
	jediService := service.NewJediService(jediRepository)
	jediController := controller.NewJediController(jediService)

	var router = gin.Default()
	router.GET("/api/jedi/:id", jediController.GetJediById)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
