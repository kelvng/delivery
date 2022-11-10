package ginnrestaurant

import (
	"awesomeProject1/common"
	"awesomeProject1/component/appctx"
	restaurantbiz "awesomeProject1/module/restaurant/biz"
	restaurantmodel "awesomeProject1/module/restaurant/model"
	restaurantstorage "awesomeProject1/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMaiDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		//go func() {
		//	defer common.AppRecover()
		//
		//	arr := []int{}
		//	log.Println(arr[0])
		//}()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
