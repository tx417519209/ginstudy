package main

import (
	"awesomeProject/src/gin-blog/pkg/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	gin.SetMode(settings.RunMode)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:      r,
		ReadTimeout:  settings.ReadTimeout,
		WriteTimeout: settings.WriteTimeout,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("ListenAndServe Error:%v\n", err)
		os.Exit(2)
	}
}
