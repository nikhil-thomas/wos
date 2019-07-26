package main


import (
	"github.com/gin-gonic/gin"
	"go-mod-1/handlers"
)


func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run()
}
