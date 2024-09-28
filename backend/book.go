package main

import "fmt"
import "rsc.io/quote"
import "github.com/gin-gonic/gin" // gin framework
import "net/http" // display http status



func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())

	//declare router
	router := gin.Default()
	// get book listing
	router.GET("/listing", func(c *gin.Context) {
		// c.JSON(http.StatusOK, listing)
		
	})
}