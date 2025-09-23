package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables only in development
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}
	}

	log.Println("Starting auth service...")

	// Initialize database connection
	log.Println("🔄 Initializing database connection...")
	if err := InitDB(); err != nil {
		log.Printf("❌ WARNING: Failed to connect to database: %v", err)
		log.Println("⚠️  Service will start without database connectivity")
		log.Println("⚠️  Database-dependent endpoints will not work until connection is established")
	} else {
		log.Println("✅ Database connection established successfully")
		defer CloseDB()
	}

	// Set Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if ginMode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize router
	router := gin.Default()

	// Initialize routes
	setupRoutes(router)

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Auth Service starting on port %s", port)
	log.Printf("Health endpoint available at: http://localhost:%s/health", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(router *gin.Engine) {
	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check - handle both root and base path
	router.GET("/health", func(c *gin.Context) {
		dbStatus := "disconnected"
		if IsDBConnected() {
			dbStatus = "connected"
		}

		c.JSON(200, gin.H{
			"status":   "ok",
			"service":  "auth-service",
			"database": dbStatus,
		})
	})

	// Health check with base path for Choreo routing
	router.GET("/auth-service/health", func(c *gin.Context) {
		dbStatus := "disconnected"
		if IsDBConnected() {
			dbStatus = "connected"
		}

		c.JSON(200, gin.H{
			"status":   "ok",
			"service":  "auth-service",
			"database": dbStatus,
		})
	})

	// Auth routes (original paths)
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", handleRegister)
		auth.POST("/login", handleLogin)
		auth.GET("/verify", handleVerify)
	}

	// Auth routes with base path for Choreo routing
	authWithBasePath := router.Group("/auth-service/api/auth")
	{
		authWithBasePath.POST("/register", handleRegister)
		authWithBasePath.POST("/login", handleLogin)
		authWithBasePath.GET("/verify", handleVerify)
	}
}
