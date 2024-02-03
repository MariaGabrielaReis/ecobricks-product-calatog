package database

import (
	"database/sql"

	"github.com/MariaGabrielaReis/imersao-full-cycle-product-catalog/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name FROM products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID, &product.ImageURL, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product

	err := pd.db.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID, &product.ImageURL, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pd *ProductDB) GetProductByCategory(categoryId string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, category_id, image_url, price FROM products WHERE category_id = ?", categoryId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID, &product.ImageURL, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}