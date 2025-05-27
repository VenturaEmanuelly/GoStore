package usecase

import (
	"store/internal/controllers"
	"store/internal/entity"
)

type Sistema struct {
	repo controllers.Repositorio
}


func (s Sistema) CalculoDeItens(consulta entity.Consulta) (entity.Consulta, error){
var total float64
	var itensCompletos []entity.Cadastro

	for _, item := range consulta.Iten {
		dbItem, err := s.repo.Get(item.Code)
		if err != nil {
			return entity.Consulta{}, err
		}
		total += dbItem.Price
		itensCompletos = append(itensCompletos, dbItem)
	}

	consulta.Iten = itensCompletos
	consulta.Total = total

	return consulta, nil
}

func NewSistemas(repo controllers.Repositorio) controllers.Sistema {
	return Sistema{repo: repo}
}