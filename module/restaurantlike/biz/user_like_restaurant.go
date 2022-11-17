package rstlikebbiz

import (
	"awesomeProject1/common"
	restaurantlikemodel "awesomeProject1/module/restaurantlike/model"
	"context"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikedCountRessStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikedCountRessStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	incStore IncLikedCountRessStore,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:    store,
		incStore: incStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
