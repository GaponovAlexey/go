package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var nWeapon = map[string]string{
	"ninja": "start",
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		wType := c.Query("type")
		wName, ok := nWeapon[wType]
		if !ok {
			c.JSON(404, gin.H{
				"status": "404",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			wType: wName,
		})
	})
	r.Run(":3000")
}
