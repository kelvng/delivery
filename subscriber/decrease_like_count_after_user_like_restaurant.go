package subscriber

import (
	"awesomeProject1/component/appctx"
	restaurantstorage "awesomeProject1/module/restaurant/storage"
	"awesomeProject1/pubsub"
	"context"
	"log"
)

func DecreaseLikeCountAfterUserLikeRestaurant(
	appCtx appctx.AppContext) consumerJob {

	return consumerJob{
		Title: "Decrease like count after user dislike restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}

func PushNotificationWhenUserDisLikeRestaurant(
	appCtx appctx.AppContext) consumerJob {

	return consumerJob{
		Title: "Push notification when user dislike restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push notification when user dislike restaurant id:", likeData.GetRestaurantId())

			return nil
		},
	}

}
