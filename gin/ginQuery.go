package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var nWeapon = map[string]interface{}{
	"ninja": "start",
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { // 3000/?type=ninja
		
    wType := c.Query("type")

		wName := nWeapon[wType]

		c.JSON(http.StatusOK, gin.H{
			wType: wName,
		})
	})
	r.Run(":3000")
}
