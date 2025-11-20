package repository

import entities "menu/src/Endpoint/Canciones/Domain/Entities"

type Icancion interface {
	GetAll()([]*entities.Canciones, error)
}