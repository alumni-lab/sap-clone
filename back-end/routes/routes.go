package routes

import "github.com/gin-gonic/gin"

func Routes() {
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
			"message": "/vendors/" + id + " deleted",
		})
	})

	router.POST("/vendors/:id/addresses", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendors/" + id + "/addresses",
		})
	})

	router.POST("/vendors/:id/addresses/:address_id", func(c *gin.Context) {
		id := c.Param("id")
		address := c.Param("address_id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendors/" + id + "/addresses/" + address,
		})
	})

	// Get customers
	router.GET("/customer-orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "/customers-orders",
		})
	})

	router.GET("/customer-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "/customer-orders/" + id,
		})
	})


	//Post customer orders

	router.POST("/customer-orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/customer-orders",
		})
	})
	router.POST("/customer-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/customer-orders/" + id,
		})
	})

	router.POST("/customer-orders/:id/delete", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/customer-orders/" + id + "/delete",
		})
	})

	// Get vendor orders
	router.GET("/vendor-orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "/vendor-orders",
		})
	})

	router.GET("/vendor-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya got got",
			"message": "/vendor-orders/" + id,
		})
	})


	//Post vendor orders

	router.POST("/vendor-orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendor-orders post",
		})
	})
	router.POST("/vendor-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendor-orders/" + id + " posted",
		})
	})

	router.POST("/vendor-orders/:id/delete", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendor-orders/" + id + "/delete",
		})
	})
	router.Run(":8000")
}