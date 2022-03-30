package ginrestaurant

import (
	"200lab/common"
	"200lab/components"
	"200lab/modules/restaurant/restaurantbiz"
	"200lab/modules/restaurant/restaurantmodel"

	"200lab/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(data))
	}
}

// type fakeCreateStore struct{}

// func (fakeCreateStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
// 	data.Id = 10
// 	return nil
// }
