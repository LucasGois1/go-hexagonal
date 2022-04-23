package app_test

import (
	"testing"

	mock_app "github.com/LucasGois1/go-hexagonal/app/mocks"

	"github.com/LucasGois1/go-hexagonal/app"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"
)

func TestProductService_Create(t *testing.T) {
	// Arrange
	controller := gomock.NewController(t)

	mockProduct := app.NewProduct()
	mockProduct.Name = "validName"
	mockProduct.Price = 10.0

	productPersistenceStub := mock_app.NewMockProductPersistenceInterface(controller)

	productPersistenceStub.EXPECT().Save(mockProduct).Return(mockProduct, nil).Times(1)

	sut := app.ProductService{
		Persistence: productPersistenceStub,
	}

	// Act
	product, err := sut.Create(mockProduct)

	// Assert
	require.Nil(t, err)
	require.Equal(t, mockProduct, product)
}

func TestProductService_Get(t *testing.T) {
	// Arrange
	controller := gomock.NewController(t)

	mockProduct := mock_app.NewMockProductInterface(controller)
	productPersistenceStub := mock_app.NewMockProductPersistenceInterface(controller)

	productPersistenceStub.EXPECT().Get("validId").Return(mockProduct, nil).Times(1)

	sut := app.ProductService{
		Persistence: productPersistenceStub,
	}

	// Act
	product, err := sut.Get("validId")

	// Assert
	require.Nil(t, err)
	require.Equal(t, mockProduct, product)
}

func TestProductService_Enable(t *testing.T) {
	// Arrange
	controller := gomock.NewController(t)

	mockProduct := mock_app.NewMockProductInterface(controller)
	mockProduct.EXPECT().Enable().Return(nil).Times(1)
	mockProduct.EXPECT().IsValid().Return(true, nil).Times(1)
	productPersistenceStub := mock_app.NewMockProductPersistenceInterface(controller)

	productPersistenceStub.EXPECT().Save(mockProduct).Return(mockProduct, nil).Times(1)

	sut := app.ProductService{
		Persistence: productPersistenceStub,
	}

	// Act
	product, err := sut.Enable(mockProduct)

	// Assert
	require.Nil(t, err)
	require.Equal(t, mockProduct, product)
}

func TestProductService_Disable(t *testing.T) {
	// Arrange
	controller := gomock.NewController(t)

	mockProduct := mock_app.NewMockProductInterface(controller)
	mockProduct.EXPECT().Disable().Return(nil).Times(1)
	mockProduct.EXPECT().IsValid().Return(true, nil).Times(1)
	productPersistenceStub := mock_app.NewMockProductPersistenceInterface(controller)

	productPersistenceStub.EXPECT().Save(mockProduct).Return(mockProduct, nil).Times(1)

	sut := app.ProductService{
		Persistence: productPersistenceStub,
	}

	// Act
	product, err := sut.Disable(mockProduct)

	// Assert
	require.Nil(t, err)
	require.Equal(t, mockProduct, product)
}
