package ginrstlike

import (
	"awesomeProject1/common"
	"awesomeProject1/component/appctx"
	rstlikebbiz "awesomeProject1/module/restaurantlike/biz"
	restaurantlikestorage "awesomeProject1/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DELETE /v1/restaurants/:id/unlike

func UserUnlikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMaiDBConnection())

		biz := rstlikebbiz.NewUserDisLikeRestaurantBiz(store)

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
