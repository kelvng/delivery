package rstlikebbiz

import (
	"awesomeProject1/component/asyncjob"
	restaurantlikemodel "awesomeProject1/module/restaurantlike/model"
	"context"
	"log"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecLikedCountRessStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store    UserDislikeRestaurantStore
	decStore DecLikedCountRessStore
}

func NewUserDisLikeRestaurantBiz(
	store UserDislikeRestaurantStore,
	decStore DecLikedCountRessStore,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store:    store,
		decStore: decStore,
	}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	// Side effect
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	return nil
}
