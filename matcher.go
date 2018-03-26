package main

import (
	"github.com/gin-gonic/gin"
)

// Data struct to hold JSON data
type Data struct {
	Hexs []string `binding:"required"`
	URL  string   `binding:"required"`
}

func main() {
	// Setup router
	r := gin.Default()

	// Simple group: v1
	paletteMatcher := r.Group("/api/palette_matcher")
	{
		paletteMatcher.POST("/match", func(c *gin.Context) {
			data := new(Data)
			err := c.BindJSON(data)
			if err != nil {
				c.AbortWithError(400, err)
				return
			}
			c.JSON(200, gin.H{
				"isMatch":   true, // Was the image matched with only colors provided
				"numColors": 5,    // How many colors exist (this can be used a verification)
			})
		})

		paletteMatcher.GET("/exampleData", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"isMatch":   true,
				"numColors": 5,
			})
		})
	}
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}

func checkMatch() {

}

func countColors() {

}
