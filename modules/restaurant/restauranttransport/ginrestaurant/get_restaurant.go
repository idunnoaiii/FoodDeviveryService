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

func GetRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)
		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(data))
	}
}
