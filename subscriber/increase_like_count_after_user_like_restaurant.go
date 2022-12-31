package subscriber

import (
	"awesomeProject1/component/appctx"
	restaurantstorage "awesomeProject1/module/restaurant/storage"
	"awesomeProject1/pubsub"
	"context"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	//GetUserId() int
}

//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicLikeRestaurant)
//
//	store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
//
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//
//}

func IncreaseLikeCountAfterUserLikeRestaurant(
	appCtx appctx.AppContext) consumerJob {

	return consumerJob{
		Title: "Increase like count after user like restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}

func PushNotificationWhenUserLikeRestaurant(
	appCtx appctx.AppContext) consumerJob {

	return consumerJob{
		Title: "Push notification when user like restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push notification when user like restaurant id:", likeData.GetRestaurantId())

			return nil
		},
	}

}
