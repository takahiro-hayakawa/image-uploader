package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")

		if err != nil {
			panic(err)
		}

		const uploadPath = "./image/"
		uploadFullPath := uploadPath + img.Filename
		err = img.Save(uploadFullPath)

		if err != nil {
			panic(err)
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"image_url": "http://localhost/image/",
		})
	})

	r.Run(":3000")
}
