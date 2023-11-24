package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/create", func(c *gin.Context) {
		u := CreateUrl{}
		c.Bind(&u)
		err := create(u)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.String(200, "Success")
	})
	r.Run("localhost:8080")
}
