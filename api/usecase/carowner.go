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

// CarOwnerUsecaseのインターフェースを作成(後に作成するHandlerのテストの際にUsecseの際し変えが見込まれるから)
type CarOwnerUsecaseIF interface {
    RegistCarOwner(owner *model.CarOwner) error
    FindByID(id int) (*model.CarOwner, error)
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
		return fmt.Errorf("免許証期限切れの為登録不可: %v", owner.LicenseExpiration)
	}

	return uc.CarOwnerRepo.Save(owner)
	//【蛇足説明】何でerror（エラー型）にuc.CarOwnerRepo.Save(owner)が返るのよ
	//これ、uc.CarOwnerRepo.Save(owner)はerrorを返す
	//
	//err = uc.CarOwnerRepo.Save(owner)
	//return err と同義（errにnilが返ろうがerrorが返ろうがエラー型に返せるでしょ）
}


// owner検索（ID）
func (uc *CarOwnerUsecase) FindByID(id int) (*model.CarOwner, error) {
    //idのバリデーション
	//一時的にCarOwnerインスタンスを作成してIsIDPositiveメソッドを実行
	tempOwner := &model.CarOuner{ID: id}
    if !tempOwner.IsIDPositive(id) {
		return nil, fmt.Errorf("IDが不正です(負の数): %v". id)
	}
    //idのOwner存在するか
    owner, err := uc.CarOwnerRepo.FindByID(id)
    
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

