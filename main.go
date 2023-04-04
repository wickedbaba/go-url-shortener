package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wickedbaba/go-url-shortener/handler"
	"github.com/wickedbaba/go-url-shortener/store"
)

var pl = fmt.Println
var pf = fmt.Printf

func errChecker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pl("the code is working")
	pf("\nfine\n")

	// Default returns an Engine instance with the Logger
	// and Recovery middleware already attached.
	r := gin.Default()

	//  endpoint should be in seperate files
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// store initialization
	store.InitializeStore()

	err := r.Run(":8080")
	errChecker(err)

}
