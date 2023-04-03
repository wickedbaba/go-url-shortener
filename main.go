package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener!",
		})
	})

	err := r.Run(":8080")

	errChecker(err)

}
