package subscriber

import (
	"context"
	"food_delivery/component/appctx"
	"food_delivery/pubsub"
	"log"
)

// func PushNotificationWhenUserLogin(appCtx appctx.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserUnLikeRestaurant)

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId) // Convert to HasRestaurantId type
// 			log.Println("Push notification to account: ", likeData.GetRestaurantId())
// 		}

//		}()
//	}
func RunPushNotificationWhenUserLogin(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push Notification",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId) // Convert to HasRestaurantId type
			log.Println("Push notification to account: ", likeData.GetRestaurantId())
			return nil
		},
	}
}
