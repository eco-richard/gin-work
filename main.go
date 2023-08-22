package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	r := gin.Default()
	fmt.Println("Okay")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	// Now we can combine gorm and gin

	r.GET("/products/:id", func(c *gin.Context) {
		productIdStr := c.Param("id")

		productId, err := strconv.Atoi(productIdStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Product ID"})
		}

		// Read
		var product Product
		db.First(&product, productId) // find with integer primary key

		c.JSON(http.StatusOK, product)
	})

	// Update

	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	err = r.Run() // running on port :8080
	if err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
