package repository

import (
	"github.com/taatolu/ParkingHub/api/domain/model"
)

type CarRepository interface {
	Save(car *model.Car) error
	GetAll() ([]*model.Car, error)
	FindByOwnerID(ownerID uint) ([]*model.Car, error)
	FindByID(id uint) (*model.Car, error)
	Update(car *model.Car) error
	Delete(id uint) error
}
