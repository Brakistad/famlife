package main
import "github.com/gin-gonic/gin"
// import "net/http"

import (
	"fmt"

	"github.com/Brakistad/famlife/gofamlife"
	"rsc.io/quote"
)

func main() {
	fmt.Println(gofamlife.Famlifer())
	fmt.Println(quote.Go())
	fmt.Println("It has begun")
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8000 (for windows "localhost:8080")

}