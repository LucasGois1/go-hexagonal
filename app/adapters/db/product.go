package db

import (
	"database/sql"

	"github.com/LucasGois1/go-hexagonal/app"

	_ "github.com/mattn/go-sqlite3"
)

type ProductPersistenceService struct {
	db *sql.DB
}

func NewProductPersistenceService(db *sql.DB) *ProductPersistenceService {
	return &ProductPersistenceService{db: db}
}

func (p *ProductPersistenceService) Get(id string) (app.ProductInterface, error) {
	var product app.Product

	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return &product, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return &product, err
	}

	return &product, nil
}

func (p *ProductPersistenceService) Save(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
