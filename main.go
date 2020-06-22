package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const DOWNLOADS_PATH = "./downloads/"

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("public/*")

	// index.html
	router.Static("/file", DOWNLOADS_PATH)

	// 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	// request
	router.POST("/upload", func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)
		targetPath := filepath.Join(DOWNLOADS_PATH, filename)
		if err := c.SaveUploadedFile(file, targetPath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully. ", file.Filename))
	})

	///router.Use(static.Serve("/file", static.LocalFile("/downloads", false)))
	///router.StaticFS("/file", http.Dir("downloads") )
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"status": http.StatusOK})
	})

	router.Run(":2077")
}
