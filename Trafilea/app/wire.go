package app

import (
	"TRAFILEA/handler"
	"TRAFILEA/usecases"
)

func DependencyInjection() *handler.Operations {

	multiplesOperation := usecases.NewType()

	serviceMultipliers := usecases.NewServiceMultipliers(multiplesOperation)

	Operations := handler.NewOperations(serviceMultipliers)

	return Operations
}
