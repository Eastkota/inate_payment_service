package main

import (
	"inapp_payment_service/repositories"
	"inapp_payment_service/resolvers"
	"inapp_payment_service/services"
	"inapp_payment_service/graph"
	"inapp_payment_service/helpers"
	"inapp_payment_service/handlers"

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := helpers.GetGormDB()
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}
	inatePaymentRepository := repositories.NewInatePaymentRepository(db)
	inatePaymentService := services.NewInatePaymentService(inatePaymentRepository)
	resolver := resolvers.NewInatePaymentResolver(inatePaymentService)

	mutationType := schema.NewMutationType(resolver)
	queryType := schema.NewQueryType(resolver)

	schema.InitSchema(queryType, mutationType)

	e := echo.New()
	e.Use(handlers.RecoverMiddleware())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	e.GET("/health", func(c echo.Context) error {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{"status": "unhealthy"})
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	e.POST("/graphql", handlers.Handler)

	go func() {
		if err := e.Start(":8103"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}

	log.Println("Server shut down gracefully")
}
