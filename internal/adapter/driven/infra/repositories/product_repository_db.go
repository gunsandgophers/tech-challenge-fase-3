package repositories

import (
	"strings"
	"tech-challenge-fase-1/internal/adapter/driven/infra/database"
	"tech-challenge-fase-1/internal/core/domain/entities"
	"tech-challenge-fase-1/internal/core/domain/errors"
)

type ProductRepositoryDB struct {
	conn database.ConnectionDB
}

func NewProductRepositoryDB(conn database.ConnectionDB) *ProductRepositoryDB {
	return &ProductRepositoryDB{conn: conn}
}

func (r *ProductRepositoryDB) CreateProduct(product *entities.Product) error {
	query := "INSERT INTO products (name, category, price, description, image ) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := r.conn.QueryRow(query, product.Name, product.Category, product.Price, product.Description, product.Image)
	err := row.Scan(&product.ID)
	if err != nil {
		if strings.Contains(err.Error(), "23514") {
			return errors.ErrInvalidCategory
		}
		return err
	}

	return nil
}

func (r *ProductRepositoryDB) FindByID(ID int) (*entities.Product, error) {
	query := "SELECT id, name, price, category, description, image FROM products WHERE id = $1"

	row := r.conn.QueryRow(query, ID)

	var product entities.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Category, &product.Description, &product.Image)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.ErrProductNotFound
		}

		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositoryDB) UpdateProduct(product *entities.Product) (*entities.Product, error) {
	query := "UPDATE products SET name = $1, category = $2, price = $3, description = $4, image = $5, updated_at = NOW() WHERE id = $6 RETURNING id"
	row := r.conn.QueryRow(query, product.Name, product.Category, product.Price, product.Description, product.Image, product.ID)
	err := row.Scan(&product.ID)
	if err != nil {
		if strings.Contains(err.Error(), "23514") {
			return nil, errors.ErrInvalidCategory
		}
		return nil, err
	}

	return product, nil
}

func (r *ProductRepositoryDB) DeleteProduct(ID int) error {
	query := "DELETE FROM products WHERE id = $1"
	err := r.conn.Exec(query, ID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.ErrProductNotFound
		}

		return err
	}

	return nil
}

func (r *ProductRepositoryDB) ListProducts() ([]*entities.Product, error) {
	query := "SELECT * FROM products"
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}

	var products []*entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category, &product.Description, &product.Image)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (r *ProductRepositoryDB) FindByCategory(category string, page, size int) ([]*entities.Product, error) {
	query := "SELECT id, name, price, category, description, image FROM products WHERE category = $1 LIMIT $2 OFFSET $3"
	rows, err := r.conn.Query(query, category, size, (page-1)*size)
	if err != nil {
		return nil, err
	}

	var products []*entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category, &product.Description, &product.Image)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
