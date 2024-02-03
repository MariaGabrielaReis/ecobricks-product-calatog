package service

import (
	"github.com/MariaGabrielaReis/ecobricks-product-catalog/internal/database"
	"github.com/MariaGabrielaReis/ecobricks-product-catalog/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(ProductDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: ProductDB}
}

func (cs *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := cs.ProductDB.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (cs *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := cs.ProductDB.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (cs *ProductService) GetProductByCategory(categoryId string) ([]*entity.Product, error) {
	products, err := cs.ProductDB.GetProductByCategory(categoryId)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (cs *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, category_id, image_url, price)
	_, err := cs.ProductDB.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
