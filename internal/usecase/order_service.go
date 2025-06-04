package usecase

import (
	"store/internal/controllers"
	"store/internal/entity"
)

type orderService struct {
	repo controllers.Repository
}

func (o orderService) CalculateOrder(order entity.Order) (entity.Order, error) {

	err := order.ValidateOrder()
	if err != nil {
		return entity.Order{}, err
	}

	var total float64
	var detailedItems []entity.Product

	for _, item := range order.Items {
		dbItem, err := o.repo.Get(item.Code)
		if err != nil {
			return entity.Order{}, err
		}

		total += dbItem.Price
		detailedItems = append(detailedItems, dbItem)
	}

	order.Items = detailedItems
	order.Total = total

	return order, nil
}

func NewOrderServices(repo controllers.Repository) controllers.OrderService {
	return orderService{repo: repo}
}
