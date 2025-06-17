package repository

import (
	"store/internal/controllers"
	"store/internal/entity"
)

type repository struct {
	repoSql controllers.RepoSql
}

func (r repository) Insert(product entity.Product) (entity.Product, error) {
	var productReturn entity.Product

	err := r.repoSql.QueryRow(`INSERT INTO product (code,name,price) VALUES ($1,$2,$3) RETURNING code, name, price`,
		[]any{product.Code, product.Name, product.Price}, &productReturn.Code, &productReturn.Name, &productReturn.Price)

	return productReturn, err

}

func (r repository) Get(product string) (entity.Product, error) {
	var productReturn entity.Product

	err := r.repoSql.QueryRow(`SELECT code, name, price FROM product WHERE code=$1`,
		[]any{product}, &productReturn.Code, &productReturn.Name, &productReturn.Price)

	return productReturn, err
}
func (r repository) Update(product entity.Product) (entity.Product, error) {
	
	current, err := r.Get(product.Code)
	if err != nil {
		return entity.Product{}, err
	}

	if product.Name == "" {
		product.Name = current.Name
	}
	if product.Price == 0 {
		product.Price = current.Price
	}

	var updated entity.Product
	err = r.repoSql.QueryRow(
		`UPDATE product SET name=$2, price=$3 WHERE code=$1 RETURNING code, name, price`,
		[]any{product.Code, product.Name, product.Price},
		&updated.Code, &updated.Name, &updated.Price,
	)
	return updated, err
}


func (r repository) Delete(code string) error {
	
	_, err := r.repoSql.Exec(`DELETE FROM product WHERE code = $1`, code)
	return err
}

func (r repository) InitSchema() error {
	_, err := r.repoSql.Exec(`
		CREATE TABLE IF NOT EXISTS product (
			code VARCHAR(50) PRIMARY KEY,
			name VARCHAR(100),
			price NUMERIC 
		)
	`)
	return err
}

func Newrepository(rep controllers.RepoSql) controllers.Repository {
	return repository{repoSql: rep}
}
