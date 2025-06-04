package usecase

import (
	"store/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validations(t *testing.T) {
	validCode := "547df4%"
	validName := "notebook"
	validPrice := 8.000

	t.Run("missing code", func(t *testing.T) {
		product := entity.Product{
			Code:  "",
			Name:  validName,
			Price: validPrice,
		}
		err := product.ValidateProduct()
		assert.ErrorContains(t, err, "Code")
		assert.ErrorContains(t, err, "required")
	})

	t.Run("missing name", func(t *testing.T) {
		product := entity.Product{
			Code:  validCode,
			Name:  "",
			Price: validPrice,
		}
		err := product.ValidateProduct()
		assert.ErrorContains(t, err, "Name")
		assert.ErrorContains(t, err, "required")
	})

	t.Run("missing price", func(t *testing.T) {
		product := entity.Product{
			Code:  validCode,
			Name:  validName,
			Price: 0,
		}
		err := product.ValidateProduct()
		assert.ErrorContains(t, err, "Price")
		assert.ErrorContains(t, err, "gt")
	})

}
