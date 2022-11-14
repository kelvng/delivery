package restaurantbiz

import (
	restaurantmodel "awesomeProject1/module/restaurant/model"
	"context"
	"errors"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store: store,
	}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldDData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldDData.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := biz.store.Delete(context, id); err != nil {
		return err
	}

	return nil
}
