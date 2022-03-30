package main

import (
	"log"
	"os"

	"200lab/components"
	"200lab/middleware"
	"200lab/modules/restaurant/restauranttransport/ginrestaurant"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	dbConnectionString := os.Getenv("DB_Connection_String")
	db, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error open database", err)
	}

	appCtx := components.NewAppContext(db)

	r := gin.Default() 

	r.Use(middleware.Recorver(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:id", ginrestaurant.UpdadeRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	r.Run()
}
