package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Get Vendors
	router.GET("/vendors", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "/vendors",
		})
	})

	router.GET("/vendors/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "/vendors/" + id,
		})
	})

	router.GET("/vendors/:id/addresses", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "/vendors/" + id + "/addresses",
		})
	})

	router.GET("/vendors/:id/addresses/:address_id", func(c *gin.Context) {
		id := c.Param("id")
		address :=  c.Param("address_id")
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "/vendors/" + id + "/addresses" + address ,
		})
	})

	//Post vendors

	router.POST("/vendors", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendors",
		})
	})
	router.POST("/vendors/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendors/" + id,
		})
	})

	router.POST("/vendors/:id/delete", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendors/"id + " deleted",
		})
	})

	router.POST("/vendors/:id/addresses", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "you in addresses" + id,
		})
	})

	router.POST("/vendors/:id/addresses/:address_id", func(c *gin.Context) {
		id := c.Param("id") + " " + c.Param("address_id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": id,
		})
	})

	// Get customers
	router.GET("/customer-orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "all customers",
		})
	})

	router.GET("/customer-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": id + " customer id",
		})
	})


	//Post customer orders

	router.POST("/customer-orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "post customer order",
		})
	})
	router.POST("/customer-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": id + " post customer order",
		})
	})

	router.POST("/customer-orders/:id/delete", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": id + " customer order deleted",
		})
	})

	// Get vendor orders
	router.GET("/vendor-orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "all vendor id",
		})
	})

	router.GET("/vendor-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": id + " vendor order id",
		})
	})


	//Post vendor orders

	router.POST("/vendor-orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "post vendor order",
		})
	})
	router.POST("/vendor-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": id + " post vendor order",
		})
	})

	router.POST("/vendor-orders/:id/delete", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": id + " vendor order deleted",
		})
	})
	router.Run(":8000")
}