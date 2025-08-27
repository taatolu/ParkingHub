package mocks

import(
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/domain/service"
)

type FakeCarOwnerRepo struct{
	SavedOwner	*model.CarOwner
	allOwners	[]*model.CarOwner
}

func (f *FakeCarOwnerRepo) Save (carOwner *model.CarOwner) error {
	//nameの入力不足バリデーション
	if !service.CarOwnerNameValidation(carOwner) {
		//CarOwnerNameValidationがFalseだったら
		return fmt.Errorf("少なくとも２つ以上のフィールドに名前を入力ください")
	}
	
	//免許証の日付不正のバリデーション
	if carOwner.IsLicenseExpired(){
		//isLicenseExpiredがtrueだったら（期限切れだったら）、、、
		return fmt.Errorf("免許証期限切れの為登録不可 %v", carOwner.LicenseExpiration)
	}

	//名前の入力条件も免許の期限もPassした場合は
	f.SavedOwner = carOwner
	return nil
}


func (f *FakeCarOwnerRepo) FindByID(id int) (*model.CarOwner, error) {
	tempOwner := &model.CarOwner{ID:id}
	if !tempOwner.IsIDPositive(){
		return nil, fmt.Errorf("IDが不正です(負の数): %v", id)
	}
    return tempOwner, nil
}

func (f *FakeCarOwnerRepo) FindByName(name string)([]*model.CarOwner, error) {
	foundOwners := []*model.CarOwner{}

	for _, owner := range f.allOwners {
		if owner.IsContainsName(name) {
			foundOwners = append(foundOwners, owner)
		}
	}
	return foundOwners, nil
}