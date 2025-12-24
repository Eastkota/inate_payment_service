package main

import (
	"inate_payment_service/repositories"
	"inate_payment_service/resolvers"
	"inate_payment_service/services"
	"inate_payment_service/graph"
	"inate_payment_service/helpers"
	"log"

	"inate_payment_service/handlers"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Add any origins you need
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	e.POST("/graphql", handlers.Handler)

	e.Logger.Fatal(e.Start(":8095"))
}



