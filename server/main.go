package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "success",
		})
	})

	r.GET("/image/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		c.File("./image/" + fileName)
	})

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
			"image_url": "http://localhost:3000/image/" + img.Filename,
		})
	})

	r.Run(":3000")
}
