package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"user-api/internal/handler"
	"user-api/internal/middleware"
	mongoRepo "user-api/internal/repository/mongodb"
	"user-api/internal/service"
	"user-api/pkg/auth"
)

func main() {
	// Get configuration from environment variables
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	jwtSecret := getEnv("JWT_SECRET", "default-secret-key")
	port := getEnv("PORT", "8080")

	log.Println("Starting server...")
	log.Println("MongoDB URI:", mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	// Test MongoDB connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	log.Println("Connected to MongoDB successfully")

	// Initialize database
	db := client.Database("userdb")

	// Initialize layers
	userRepo := mongoRepo.NewMongoUserRepository(db)
	jwtService := auth.NewJWTService(jwtSecret)
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

	log.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Helper function to get environment variables with default values
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
