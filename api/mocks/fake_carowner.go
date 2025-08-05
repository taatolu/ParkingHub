package mocks

import(
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/domain/service"
)

type FakeCarOwnerRepo struct{
	SavedOwner	*model.CarOwner
	TargetOwners    []*model.CarOwner
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
    //idのバリデーション
    if id<0{
        return nil, fmt.Errorf("IDは負の値は不可: ID=%v", id)
    }
    
    //FakeCarOwnerRepoのTargetOwnersから一致するOwnerを検索
    //FakeTestではTargetOwnersに複数のownerサンプルを登録する
    owner, err := f.TargetOwners.FindByID(id)
    if err != nil{
        //errorの場合
        return nil, err
    }
    
    if owner == nil {
        //ownerのデータが存在しない場合
        return nil, fmt.Errorf("IDに対応するownerが存在しません: ID=%v", id)
    }
    
    return owner, nil
}