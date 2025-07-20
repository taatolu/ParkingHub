package usecase

import (
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/domain/repository"
	"github.com/taatolu/ParkingHub/api/domain/service"
)

// ユースケース構造体の作成
type CarOwnerUsecase struct {
	CarOwnerRepo repository.CarOwnerRepository
}

// owner登録処理
func (uc *CarOwnerUsecase) RegistCarOwner(owner *model.CarOwner) error {
	//入力漏れの確認
	if !service.CarOwnerNameValidation(owner){
		//CarOwnerNameValidationがFalseだったら
		return fmt.Errorf("少なくとも２つ以上のフィールドに名前を入力ください")
	}

	//免許証期限が切れている場合(若しくはIsLicenseExpiredが入力されていない場合)、エラーを返す
	if owner.IsLicenseExpired() {
		//IsLicenseExpiredがtrueだったら（期限切れだったら）、、、
		return fmt.Errorf("免許証期限切れの為登録不可 %v", owner.LicenseExpiration)
	}

	return uc.CarOwnerRepo.Save(owner)
}
