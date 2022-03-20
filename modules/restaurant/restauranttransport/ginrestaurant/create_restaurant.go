package ginrestaurant

import (
	"200lab/modules/restaurant/restaurantbiz"
	"200lab/modules/restaurant/restaurantmodel"
	"200lab/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CrateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, data)
	}
}
