package ginuser

import (
	"awesomeProject1/common"
	"awesomeProject1/component/appctx"
	"awesomeProject1/component/hasher"
	"awesomeProject1/component/tokenprovider/jwt"
	userbiz "awesomeProject1/module/user/biz"
	usermodel "awesomeProject1/module/user/model"
	userstore "awesomeProject1/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMaiDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, md5, tokenProvider, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
