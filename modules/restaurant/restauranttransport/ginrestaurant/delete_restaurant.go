package ginrestaurant

import (
	"200lab/common"
	"200lab/components"
	"200lab/modules/restaurant/restaurantbiz"
	"200lab/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func (c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id")); 
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		
		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(true))

	}
}