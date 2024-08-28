package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// const (
// 	HOST     = "localhost"
// 	PORT     = 5432
// 	USER     = "postgres"
// 	PASSWORD = "Asdfghjkl"
// 	DBNAME   = "lelangkuy2"
// )

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load(".env")

	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("Success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	// db.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()
	router.GET("/", index)

	router.Run(":" + os.Getenv("PORT"))
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "SELAMAT DATANG DI KASIR ONLINE BY ABDULATIF")
}
