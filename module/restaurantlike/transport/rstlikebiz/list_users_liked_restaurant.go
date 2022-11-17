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

func ListUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMaiDBConnection())
		biz := rstlikebbiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
