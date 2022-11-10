package middleware

import (
	"awesomeProject1/common"
	"awesomeProject1/component/appctx"
	"errors"
	"github.com/gin-gonic/gin"
)

func RoleRRequired(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {

	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowRoles {
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrNoPermission(errors.New("invalid role user")))
		}

		c.Next()
	}
}
