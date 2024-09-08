package main

import (
	"21_clean-architecture/internal/domain/model"
	"log"
	"net/http"
	"project/internal/infrastructure/db"
	"project/internal/interfaces/http"
	"project/internal/usecase"

	"github.com/go-chi/chi"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&model.User{})

	userRepo := db.NewUserRepository(database)
	userUseCase := &usecase.UserUseCase{UserRepo: userRepo}

	userHandler := http.NewUserHandler(userUseCase)

	router := chi.NewRouter()
	router.Post("/users", userHandler.CreateUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
