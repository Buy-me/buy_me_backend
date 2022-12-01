package subscriber

import (
	"context"

	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"food_delivery/pubsub"
)

func DecreaseLikeCountAfterUserUnLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserUnLikeRestaurant)
	store := restaurantstorage.NewSQlStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId) // Convert to HasRestaurantId type
			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}

func RunDecreaseLikeCountAfterUserUnLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
