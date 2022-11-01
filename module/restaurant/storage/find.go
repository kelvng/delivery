package restaurantstorage

import (
	restaurantmodel "awesomeProject1/module/restaurant/model"
	"context"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {

	var data restaurantmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
