package bizrestaurant

import (
	"context"
	"errors"
	restaurantmodel "go_project/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}
type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewdeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data has been deleted")
	}

	if err := biz.store.Delete(context, id); err != nil {
		return err
	}

	return nil
}
