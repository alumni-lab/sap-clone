package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	// Item Routes

	r.GET("/items", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get request made to /items",
		})
	})

	r.GET("/items/:id", func(c *gin.Context) {
		item := c.Param("id")
		c.JSON(200, gin.H{
			"message": "get request made to /items/:id",
			"item_id": item,
		})
	})

	r.POST("/items", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post request made to /items",
		})
	})

	r.POST("/items/:id", func(c *gin.Context) {
		item := c.Param("id")
		c.JSON(200, gin.H{
			"message": "post request made to items/:id",
			"item_id": item,
		})
	})

	r.POST("/items/:id/delete", func(c *gin.Context) {
		item := c.Param("id")
		c.JSON(200, gin.H{
			"message": "post request made to items/:id/delete",
			"item_id": item,
		})
	})

	// User Routes

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get request made to /users",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		user := c.Param("id")
		c.JSON(200, gin.H{
			"message": "get request made to /users/:id",
			"user_id": user,
		})
	})

	r.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post request /users",
		})
	})

	r.POST("/users/:id", func(c *gin.Context) {
		user := c.Param("id")
		c.JSON(200, gin.H{
			"message": "post request made to /users/:id",
			"user_id": user,
		})
	})

	r.POST("/users/:id/delete", func(c *gin.Context) {
		user := c.Param("id")
		c.JSON(200, gin.H{
			"message": "post request made /users/:id/delete",
			"user_id": user,
		})
	})

	// Customer Routes

	r.GET("/customers", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get request made to /customers",
		})
	})

	r.GET("/customers/:id", func(c *gin.Context) {
		customer := c.Param("id")
		c.JSON(200, gin.H{
			"message":     "get request made to /customers/:id",
			"customer_id": customer,
		})
	})

	r.GET("/customers/:id/addresses", func(c *gin.Context) {
		customer := c.Param("id")
		c.JSON(200, gin.H{
			"message":     "get request made to ",
			"customer_id": customer,
		})
	})

	r.GET("/customers/:id/addresses/:address_id", func(c *gin.Context) {
		customer := c.Param("id")
		address := c.Param("address_id")
		c.JSON(200, gin.H{
			"message":     "get request made to /customers/:id/addresses/:address_id",
			"customer_id": customer,
			"address_id":  address,
		})
	})

	r.POST("/customers", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post request made to /customers",
		})
	})

	r.POST("/customers/:id", func(c *gin.Context) {
		customer := c.Param("id")
		c.JSON(200, gin.H{
			"message":     "post request made to /customers/:id",
			"customer_id": customer,
		})
	})

	r.POST("/customers/:id/delete", func(c *gin.Context) {
		customer := c.Param("id")
		c.JSON(200, gin.H{
			"message":     "post request made to /customers/:id/delete",
			"customer_id": customer,
		})
	})

	r.POST("/customers/:id/addresses/:address_id", func(c *gin.Context) {
		customer := c.Param("id")
		address := c.Param("address_id") 
			"message":     "post request made to /customers/:id/addresses/:address_id",
			"customer_id": customer,
			"address": address,
		})
	})

	r.POST("/customers/:id/addresses/:address_id/delete", func(c *gin.Context) {
		customer := c.Param("id")
		address := c.Param("address_id")
		c.JSON(200, gin.H{
			"message":     "post request made to /customers/:id/addresses/:address_id/delete",
			"customer_id": customer,
			"address": address,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
