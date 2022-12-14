package ginrstlike

import (
	"awesomeProject1/common"
	"awesomeProject1/component/appctx"
	rstlikebbiz "awesomeProject1/module/restaurantlike/biz"
	restaurantlikemodel "awesomeProject1/module/restaurantlike/model"
	restaurantlikestorage "awesomeProject1/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMaiDBConnection())
		//incStore := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := rstlikebbiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
