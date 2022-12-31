package rstlikebbiz

import (
	"awesomeProject1/common"
	restaurantlikemodel "awesomeProject1/module/restaurantlike/model"
	"awesomeProject1/pubsub"
	"context"
	"log"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

//type DecLikedCountResStore interface {
//	DecreaseLikeCount(ctx context.Context, id int) error
//}

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
	//decStore DecLikedCountResStore
	ps pubsub.Pubsub
}

func NewUserDisLikeRestaurantBiz(
	store UserDislikeRestaurantStore,
	//decStore DecLikedCountResStore,
	ps pubsub.Pubsub,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store: store,
		//decStore: decStore,
		ps: ps,
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

	if err := biz.ps.Publish(ctx, common.TopicDislikeRestaurant,
		pubsub.NewMassage(&restaurantlikemodel.Like{RestaurantId: restaurantId})); err != nil {
		log.Println(err)
	}

	//// Side effect
	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	//})
	//
	//if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	//	log.Println(err)
	//}

	return nil
}
