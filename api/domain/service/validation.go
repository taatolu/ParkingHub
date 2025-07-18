package service

import(
	"github.com/taatolu/ParkingHub/api/domain/model"
)

//CarOwnerの名前（FirstName,MiddleName,LastName)の中に１つしか入力されていない場合-->false
func CarOwnerNameValidation (owner *model.CarOwner) bool {
	var emptyNameFields []string
	if owner.FirstName == "" { 
		emptyNameFields = append(emptyNameFields, "FirstName")
	 }
	if owner.MiddleName == "" { 
		emptyNameFields = append(emptyNameFields, "MiddleName")
	}
	if owner.LastName == "" { 
		emptyNameFields = append(emptyNameFields, "LastName")
	}
	if len(emptyNameFields) >= 2{
		return false
	}

	return true
}