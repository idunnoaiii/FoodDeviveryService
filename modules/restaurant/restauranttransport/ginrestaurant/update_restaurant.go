package ginrestaurant

import (
	"200lab/common"
	"200lab/components"
	"200lab/modules/restaurant/restaurantbiz"
	"200lab/modules/restaurant/restaurantmodel"
	"200lab/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdadeRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return 
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return 
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return 
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(true))
 
	}
}