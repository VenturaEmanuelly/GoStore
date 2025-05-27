package repositorio

import (
	"store/internal/controllers"
	"store/internal/entity"
)

type repositorio struct {
	repoSql controllers.RepoSql
}

func (r repositorio) Insert(cadastro entity.Cadastro) (entity.Cadastro, error) {
	var cadastroReturn entity.Cadastro

	err := r.repoSql.QueryRow(`INSERT INTO cadastro (code,iten,price) VALUES ($1,$2,$3) RETURNING code, iten, price`,
		[]any{cadastro.Code, cadastro.Iten, cadastro.Price}, &cadastroReturn.Code, &cadastroReturn.Iten, &cadastroReturn.Price)

	return cadastroReturn, err

}

func (r repositorio) Get(cadastro string) (entity.Cadastro, error) {
	var cadastroReturn entity.Cadastro

	err := r.repoSql.QueryRow(`SELECT code, iten, price FROM cadastro WHERE code=$1`, []any{cadastro}, &cadastroReturn.Code, &cadastroReturn.Iten, &cadastroReturn.Price)

	return cadastroReturn, err
}

func (r repositorio) InitSchema() error {
	_, err := r.repoSql.Exec(`
		CREATE TABLE IF NOT EXISTS cadastro (
			code VARCHAR(50) PRIMARY KEY,
			iten VARCHAR(100),
			price NUMERIC 
		)
	`)
	return err
}

func NewRepositorio(rep controllers.RepoSql) controllers.Repositorio {
	return repositorio{repoSql: rep}
}
