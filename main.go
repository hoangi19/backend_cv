package main

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

func contact(c *gin.Context) {
	var name, _ = c.GetPostForm("name")
	var email, _ = c.GetPostForm("email")
	var phone, _ = c.GetPostForm("phone")
	var website, _ = c.GetPostForm("website")
	// c.BindJSON(&data_contact)
	fmt.Print(name)
	fmt.Print(email)
	fmt.Print(phone)
	fmt.Print(website)
	c.JSON(http.StatusCreated, gin.H{"message": "Ok"})
}

func downloadCV(c *gin.Context) {
	c.Writer.Header().Set("Content-type", "application/octet-stream")
	c.FileAttachment("./Hoangi19_eng.pdf", "Hoangi19_eng.pdf")
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": http.StatusOK,
		})
	})

	r.GET("/api/download_cv", downloadCV)
	r.POST("/api/contact", contact)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
