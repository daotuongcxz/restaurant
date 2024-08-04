package bizrestaurant

import (
	"context"
	restaurantmodel "go_project/module/restaurant/model"
)

type FindRestaurantStore interface {
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type FindRestaurantBiz struct {
	store FindRestaurantStore
}

func NewFindRestaurantBiz(store FindRestaurantStore) *FindRestaurantBiz {
	return &FindRestaurantBiz{store: store}
}

func (biz *FindRestaurantBiz) FindRestaurant(context context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}
