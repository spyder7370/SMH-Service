package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spyder7370/SMH-Service/connections"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseConnectionUrl := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))
	fmt.Println("making postgress connection")
	db, err2 := sql.Open("postgres", databaseConnectionUrl)
	if err2 != nil {
		panic(err)
	}
	defer db.Close()
	connections.DB = db
	fmt.Println("postgress connection completed")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println(databaseConnectionUrl)
	fmt.Println("Hello")
	httpServer := gin.Default()
	httpServer.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	httpServer.Run()
}
