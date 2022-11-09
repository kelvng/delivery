package userbiz

import (
	"awesomeProject1/common"
	"awesomeProject1/component/tokenprovider"
	usermodel "awesomeProject1/module/user/model"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	//appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, hasher Hasher, tokenProvider tokenprovider.Provider, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		hasher:        hasher,
		tokenProvider: tokenProvider,
		expiry:        expiry,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
