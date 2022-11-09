package ginuser

import (
	"awesomeProject1/common"
	"awesomeProject1/component/appctx"
	"awesomeProject1/component/hasher"
	userbiz "awesomeProject1/module/user/biz"
	usermodel "awesomeProject1/module/user/model"
	userstore "awesomeProject1/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appctx.GetMaiDBConnection()

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c, &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
