package main

import (
	"fmt"
	"log"
	"new/database"
	"new/database/creatdb"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")
	
	err := database.InitlizeDb("./myDb.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDb()

	if err := creatdb.Creatdb(); err != nil {
		log.Fatalf("Failed to create database schema: %v", err)
	}

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("assets/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		isValid, role, err := database.GetLogin(username, password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		if !isValid {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(200, gin.H{
			"message": "Login successful",
			"role": role,
			"redirect": "/home",
		})
	})

	r.Run()
}
