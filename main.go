package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", getHome)
	router.GET("/article/:title", getArtile)

	router.Run()
}
func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "halo kamu",
	})
}

func getArtile(c *gin.Context) {
	title := c.Param("title")
	c.JSON(200, gin.H{
		"massage": title,
	})
}
