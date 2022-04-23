package app_test

import (
	"testing"

	"github.com/LucasGois1/go-hexagonal/app"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProductEnable(t *testing.T) {
	product := app.Product{}

	product.Name = "Product"
	product.Status = app.DISABLED
	product.Price = 10.0

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "Price must be greater than 0 to enable product", err.Error())
}

func TestProductDisable(t *testing.T) {
	product := app.Product{}

	product.Name = "Product"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()
	require.Equal(t, "Price must be less or equal 0 to disable product", err.Error())
}

func TestIsValid(t *testing.T) {
	product := app.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Product"
	product.Status = app.DISABLED
	product.Price = 10

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "INVALID"

	_, err = product.IsValid()

	require.Equal(t, "Status must be enable or disable", err.Error())

	product.Status = app.ENABLED
	product.Price = -10

	_, err = product.IsValid()

	require.Equal(t, "Price must be greater or equal 0", err.Error())

}
