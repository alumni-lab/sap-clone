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
			"message": "/vendor-orders",
		})
	})
	router.POST("/vendor-orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendor-orders/" + id,
		})
	})

	router.POST("/vendor-orders/:id/delete", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  "ya posted",
			"message": "/vendor-orders/" + id + "/delete",
		})
	})
  
	// Item Routes

	router.GET("/items", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get request made to /items",
		})
	})

	router.GET("/items/:id", func(c *gin.Context) {
		item := c.Param("id")
		c.JSON(200, gin.H{
			"message": "get request made to /items/:id",
			"item_id": item,
		})
	})

	router.POST("/items", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post request made to /items",
		})
	})

	router.POST("/items/:id", func(c *gin.Context) {
		item := c.Param("id")
		c.JSON(200, gin.H{
			"message": "post request made to items/:id",
			"item_id": item,
		})
	})

	router.POST("/items/:id/delete", func(c *gin.Context) {
		item := c.Param("id")
		c.JSON(200, gin.H{
			"message": "post request made to items/:id/delete",
			"item_id": item,
		})
	})

	// User Routes

	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get request made to /users",
		})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		user := c.Param("id")
		c.JSON(200, gin.H{
			"message": "get request made to /users/:id",
			"user_id": user,
		})
	})

	router.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post request /users",
		})
	})

	router.POST("/users/:id", func(c *gin.Context) {
		user := c.Param("id")
		c.JSON(200, gin.H{
			"message": "post request made to /users/:id",
			"user_id": user,
		})
	})

	router.POST("/users/:id/delete", func(c *gin.Context) {
		user := c.Param("id")
		c.JSON(200, gin.H{
			"message": "post request made /users/:id/delete",
			"user_id": user,
		})
	})

	// Customer Routes

	router.GET("/customers", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get request made to /customers",
		})
	})

	router.GET("/customers/:id", func(c *gin.Context) {
		customer := c.Param("id")
		c.JSON(200, gin.H{
			"message":     "get request made to /customers/:id",
			"customer_id": customer,
		})
	})

	router.GET("/customers/:id/addresses", func(c *gin.Context) {
		customer := c.Param("id")
		c.JSON(200, gin.H{
			"message":     "get request made to ",
			"customer_id": customer,
		})
	})

	router.GET("/customers/:id/addresses/:address_id", func(c *gin.Context) {
		customer := c.Param("id")
		address := c.Param("address_id")
		c.JSON(200, gin.H{
			"message":     "get request made to /customers/:id/addresses/:address_id",
			"customer_id": customer,
			"address_id":  address,
		})
	})

	router.POST("/customers", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post request made to /customers",
		})
	})

	router.POST("/customers/:id", func(c *gin.Context) {
		customer := c.Param("id")
		c.JSON(200, gin.H{
			"message":     "post request made to /customers/:id",
			"customer_id": customer,
		})
	})

	router.POST("/customers/:id/delete", func(c *gin.Context) {
		customer := c.Param("id")
		c.JSON(200, gin.H{
			"message":     "post request made to /customers/:id/delete",
			"customer_id": customer,
		})
	})

	router.POST("/customers/:id/addresses/:address_id", func(c *gin.Context) {
		customer := c.Param("id")
		address := c.Param("address_id")
		c.JSON(200, gin.H{ 
			"message":     "post request made to /customers/:id/addresses/:address_id",
			"customer_id": customer,
			"address": address,
		})
	})

	router.POST("/customers/:id/addresses/:address_id/delete", func(c *gin.Context) {
		customer := c.Param("id")
		address := c.Param("address_id")
		c.JSON(200, gin.H{
			"message":     "post request made to /customers/:id/addresses/:address_id/delete",
			"customer_id": customer,
			"address": address,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
