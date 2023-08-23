package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gusetiawn/BPJS-Test/internal/routers"
)

func main() {
	router := gin.Default()

	routers.InitRouter(router)

	router.Run(":8080")
}
