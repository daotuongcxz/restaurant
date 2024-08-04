package bizrestaurant

import (
	"context"
	"go_project/common"
	restaurantmodel "go_project/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(context context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type ListRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *ListRestaurantBiz {
	return &ListRestaurantBiz{store: store}
}

func (biz *ListRestaurantBiz) ListRestaurant(context context.Context, filter *restaurantmodel.Filter, paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
