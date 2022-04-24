package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/LucasGois1/go-hexagonal/app"
	"github.com/LucasGois1/go-hexagonal/app/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable()
	createProduct()
}

func createTable() {
	table := `CREATE TABLE products ("id" string PRIMARY KEY, "name" string, "price" float, "status" string)`

	stmt, err := Db.Prepare(table)

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
}

func createProduct() {
	stmt, err := Db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec("1", "Product 1", 1.99, "enabled")

	if err != nil {
		log.Fatal(err)
	}
}

func TestProductPersistenceService_Get(t *testing.T) {
	setup()
	defer Db.Close()

	productService := db.NewProductPersistenceService(Db)

	product, err := productService.Get("1")

	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 1.99, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
	require.Equal(t, "1", product.GetId())
}

func TestProductPersistenceService_Save(t *testing.T) {
	setup()
	defer Db.Close()

	productService := db.NewProductPersistenceService(Db)

	product := app.NewProduct()
	product.Name = "Product 2"
	product.Price = 2.99
	product.Status = "enabled"

	result, err := productService.Save(product)

	fmt.Println(reflect.TypeOf(result.GetId()))

	require.Nil(t, err)
	require.Equal(t, "Product 2", result.GetName())
	require.Equal(t, 2.99, result.GetPrice())
	require.Equal(t, "enabled", result.GetStatus())
	require.Len(t, result.GetId(), 36)
}
