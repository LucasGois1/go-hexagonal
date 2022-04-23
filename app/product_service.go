package app

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Create(product ProductInterface) (ProductInterface, error) {
	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}

	product, err = s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return product, nil
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}

	product.Enable()

	product, err = s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return product, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}

	product.Disable()

	product, err = s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return product, nil
}

// func (s *ProductService) GetAll() ([]ProductInterface, error) {
// 	// return s.Persistence.GetAllProducts()
// }
