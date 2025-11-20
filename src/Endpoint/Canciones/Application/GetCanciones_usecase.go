package application

import (
	entities "menu/src/Endpoint/Canciones/Domain/Entities"
	repository "menu/src/Endpoint/Canciones/Domain/Repository"
)

type GetAllCancionesUseCase struct {
	usecase repository.Icancion
}

func NewGetAllCanciones(app repository.Icancion)*GetAllCancionesUseCase{
	return &GetAllCancionesUseCase{
		usecase: app,
	}
}

func ( g *GetAllCancionesUseCase) Execute() ([]*entities.Canciones, error){
	draws, err := g.usecase.GetAll()
	if err != nil {
		return nil, err
	}

	if draws == nil {
		return []*entities.Canciones{}, nil
	}

	return draws, nil
}