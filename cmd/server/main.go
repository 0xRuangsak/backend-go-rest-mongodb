package main

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"user-api/internal/handler"
	"user-api/internal/middleware"
	mongoRepo "user-api/internal/repository/mongodb"
	"user-api/internal/service"
	"user-api/pkg/auth"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	// Initialize database
	db := client.Database("userdb")

	// Initialize layers
	userRepo := mongoRepo.NewMongoUserRepository(db)
	jwtService := auth.NewJWTService("your-secret-key-here") // In production, use environment variable
	userService := service.NewUserService(userRepo, jwtService)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(userService)
	userHandler := handler.NewUserHandler(userService)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	// Routes
	http.HandleFunc("/auth/register", authHandler.Register)
	http.HandleFunc("/auth/login", authHandler.Login)

	// Protected routes
	http.HandleFunc("/users", authMiddleware.RequireAuth(userHandler.GetAllUsers))
	http.HandleFunc("/users/", authMiddleware.RequireAuth(userHandler.GetUserByID))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
