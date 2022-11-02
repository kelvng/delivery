package restaurantstorage

import (
	"awesomeProject1/common"
	restaurantmodel "awesomeProject1/module/restaurant/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	// where is db?????
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
