package restaurantlikebiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/restaurantlikemodel"
	"food_delivery/pubsub"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	// incStore IncreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	// incStore IncreaseLikeCountStore,
	pubsub pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store /*incStore: incStore,*/, pubsub: pubsub}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.RestaurantLike,
) error {
	// Find, if present, return "already like"

	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Send Message
	if err := biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	//// side effect
	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	// })

	// if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	// New solution: Use pub/sub

	return nil
}
