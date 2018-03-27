package main

import (
	"fmt"
	image "image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	s "strings"

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
			checkMatch(data)
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

func checkMatch(info *Data) {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	image := downloadData(info.URL)

	fmt.Println(colorAt(image, 0, 8))
	fmt.Println(colorAt(image, 0, 9))
}

func countColors() {

}

func downloadData(url string) image.Image {
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	//open a file for writing
	path := path.Base(url)
	writeFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(writeFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	writeFile.Close()
	fmt.Println("Downloaded and saved!")

	reader, _ := os.Open(path)
	defer reader.Close()
	im, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}

	return im
}

func colorAt(img image.Image, x int, y int) string {
	rU, gU, bU, aU := img.At(y, x).RGBA()
	r, g, b, _ := int(rU/257), int(gU/257), int(bU/257), int(aU/257)
	hex := fmt.Sprintf("%x", r) + fmt.Sprintf("%x", g) + fmt.Sprintf("%x", b)
	return s.ToUpper(hex)
}
