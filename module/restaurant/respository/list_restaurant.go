package restaurantrepo

import (
	"awesomeProject1/common"
	restaurantmodel "awesomeProject1/module/restaurant/model"
	"context"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]restaurantmodel.Restaurant, error)
}

//type LikeRestaurantStore interface {
//	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
//}

type listRestaurantRepo struct {
	store ListRestaurantStore
	//likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store: store,
	}
}

func (biz *listRestaurantRepo) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	//ids := make([]int, len(result))
	//
	//for i := range ids {
	//	ids[i] = result[i].Id
	//}

	//
	//likeMap, err := biz.likeStore.GetRestaurantLike(context, ids)
	//
	//if err != nil {
	//	log.Println(err)
	//	return result, nil
	//}
	//
	//for i, item := range result {
	//	result[i].LikedCount = likeMap[item.Id]
	//}

	return result, nil
}
