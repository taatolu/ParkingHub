package usecase

import (
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/domain/repository"
)

// ユースケース構造体の作成
type CarOwnerUsecase struct {
	CarOwnerRepo repository.CarOwnerRepository
}

// owner登録処理
func (uc *CarOwnerUsecase) RegistCarOwner(owner *model.CarOwner) error {
	//入力漏れの確認
	var nameSection []string

	if owner.FirstName == "" {
		nameSection = append(nameSection, "FirstName")
	}
	if owner.MiddleName == "" {
		nameSection = append(nameSection, "MiddleName")
	}
	if owner.LastName == "" {
		nameSection = append(nameSection, "LastName")
	}

	//不足が2項目以上ある場合はerrorを返す
	if len(nameSection) >= 2 {
		return fmt.Errorf("少なくとも2つの枠に入力が必要です(入力が無いのは次の2項目): %v", nameSection)
	}

	//免許証期限が切れている場合(若しくはIsLicenseExpiredが入力されていない場合)、エラーを返す
	isLicenseExpired := owner.IsLicenseExpired()
	if isLicenseExpired {
		//isLicenseExpiredがtrueだったら（期限切れだったら）、、、
		return fmt.Errorf("免許証期限切れの為登録不可 %v", owner.LicenseExpiration)
	}

	return uc.CarOwnerRepo.Save(owner)
}
