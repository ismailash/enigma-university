package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eulbyvan/enigma-university/controller"
	"github.com/eulbyvan/enigma-university/repository"
	"github.com/eulbyvan/enigma-university/usecase"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	// run application

	// db connection
	host := "localhost"
	port := "5433"
	user := "postgres"
	password := "postgres"
	dbName := "lms_university"
	driver := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open(driver, dsn)

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	} else {
		log.Printf("KONEKSI DATABASE >>> %v", db.Ping())
	}

	defer db.Close()

	// initialize repository
	userRepo := repository.NewUserRepository(db)

	// initialize usecase
	userUsecase := usecase.NewUserUseCase(userRepo)

	// initialize controller
	userCtrl := controller.NewUserController(userUsecase)

	// create gin router
	router := gin.Default()

	router.GET("/users", userCtrl.FindById)

	// run server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start Gin server: %v", err)
	}
}
