package usecase

import (
	"store/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) Get(code string) (entity.Product, error) {
	args := m.Called(code)
	return args.Get(0).(entity.Product), args.Error(1)
}

func (m *mockRepo) Insert(product entity.Product) (entity.Product, error) {
	args := m.Called(product)
	return args.Get(0).(entity.Product), args.Error(1)
}

func (m *mockRepo) Update(product entity.Product) (entity.Product, error) {
	args := m.Called(product)
	return args.Get(0).(entity.Product), args.Error(1)
}

func (m *mockRepo) Delete(product string) error {
	args := m.Called(product)
	return args.Error(1)
}

func (m *mockRepo) InitSchema() error {
	args := m.Called()
	return args.Error(0)
}

func TestCalculateOrder(t *testing.T) {
	repo := new(mockRepo)

	service := NewOrderServices(repo)

	product := entity.Product{
		Code:  "123",
		Name:  "Notebook",
		Price: 2000,
	}
	repo.On("Get", "123").Return(product, nil)

	orderInput := entity.Order{
		Items: []entity.Product{
			{Code: "123"},
		},
	}

	result, err := service.CalculateOrder(orderInput)

	assert.NoError(t, err)
	assert.Equal(t, 2000.0, result.Total)
	assert.Len(t, result.Items, 1)
	assert.Equal(t, "Notebook", result.Items[0].Name)

	repo.AssertExpectations(t)
}
func TestOrderValidation(t *testing.T) {
	t.Run("missing items", func(t *testing.T) {
		order := entity.Order{
			Items: []entity.Product{},
			Total: 100,
		}
		err := order.ValidateOrder()
		assert.ErrorContains(t, err, "at least one item")
	})

	t.Run("valid order", func(t *testing.T) {
		order := entity.Order{
			Items: []entity.Product{
				{
					Code:  "123",
					Name:  "Notebook",
					Price: 2000,
				},
			},
			Total: 2000,
		}
		err := order.ValidateOrder()
		assert.NoError(t, err)
	})
}
