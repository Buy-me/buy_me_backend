package restaurantlikebiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/restaurantlikemodel"
	"food_delivery/pubsub"
	"log"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeRestaurantBiz struct {
	store UserUnlikeRestaurantStore
	// decStore DecreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserUnlikeRestaurantBiz(
	store UserUnlikeRestaurantStore,
	// decStore DecreaseLikeCountStore,
	pubsub pubsub.Pubsub,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store: store,
		// decStore: decStore,
		pubsub: pubsub,
	}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {
	// Find, if present, if not present, return error "you have not like this restaurant"

	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	// Send Message
	if err := biz.pubsub.Publish(ctx, common.TopicUserUnLikeRestaurant, pubsub.NewMessage(
		&restaurantlikemodel.RestaurantLike{
			RestaurantId: restaurantId,
		})); err != nil {
		log.Println(err)
	}

	// side effect:
	// vì nó không ảnh hưởng đến flow chính trong api list restaurant
	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	// })

	// if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	return nil
}
