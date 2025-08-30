package repository

import (
	"github.com/taatolu/ParkingHub/api/domain/model"
)

type CarOwnerRepository interface {
	Save(carOwner *model.CarOwner) error
	FindByID(id int) (*model.CarOwner, error)
	FindByName(name string)([]*model.CarOwner, error)    //nameに一致する社員の一覧取得
	//Update(carOwner *model.CarOwner)error
	//Delete(id int)error
}
