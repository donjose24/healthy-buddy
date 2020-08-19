package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/api/handler"
	"github.com/jmramos02/healthy-buddy/internal/api/middleware"
	"github.com/jmramos02/healthy-buddy/internal/database"
	"github.com/jmramos02/healthy-buddy/internal/model"
)

func Initialize() *gin.Engine {
	db := database.Initialize()
	db.AutoMigrate(&model.User{}, &model.Customer{}, &model.Dietitian{}, &model.MealPlan{}, &model.MealPlanEntry{})
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.SetContext())
	router.Use(middleware.SetDatabase(db))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/register/customer", handler.RegisterCustomer)
	router.POST("register/dietitian", handler.RegisterDietitian)
	router.POST("/login", handler.Login)
	router.GET("/dietitians", handler.ListDietitians)

	customerRoutes := router.Group("/customer")
	{
		customerRoutes.Use(middleware.ValidateUserSession("customer"))
		customerRoutes.GET("", handler.GetCustomerDashboard)
	}

	dietitianRoutes := router.Group("/dietitian")
	{
		dietitianRoutes.Use(middleware.ValidateUserSession("dietitian"))
		dietitianRoutes.GET("", handler.GetDietitianDashboard)
		dietitianRoutes.POST("/meal", handler.CreateMealPlan)
	}

	return router
}
