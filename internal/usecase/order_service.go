package usecase

import (
	"fmt"
	"store/internal/controllers"
	"store/internal/entity"

	"github.com/go-playground/validator/v10"
)

type orderService struct {
	repo controllers.Repository
}

func (o orderService) CalculateOrder(order entity.Order) (entity.Order, error) {

	err := o.validateOrder(order)
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


func (o orderService) validateOrder(order entity.Order) error {
	validate := validator.New()
	
	err := validate.Struct(order)
	if err != nil {
		return err
	}

	
	if len(order.Items) == 0 {
		return fmt.Errorf("it is necessary to inform at least one item")
	}

	return nil
}

func NewOrderServices(repo controllers.Repository) controllers.OrderService {
	return orderService{repo: repo}
}
