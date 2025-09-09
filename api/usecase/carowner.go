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
    FindByID(id uint) (*model.CarOwner, error)
	FindByName(name string) ([]*model.CarOwner, error)
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
func (uc *CarOwnerUsecase) FindByID(id uint) (*model.CarOwner, error) {
    //idのバリデーション
	//一時的にCarOwnerインスタンスを作成してIsIDPositiveメソッドを実行
	tempOwner := &model.CarOwner{ID: id}
    if !tempOwner.IsIDPositive() {
		return nil, fmt.Errorf("IDが不正です(負の数): %v", id)
	}
    // Owner検索
    owner, err := uc.CarOwnerRepo.FindByID(id)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        // レコードが無い場合は分かりやすい日本語エラー
        return nil, fmt.Errorf("ID=%v に対応するオーナーは存在しません", id)
    }
    if err != nil {
        // その他のエラーはそのまま返却
        return nil, err
    }
    // 正常取得
    return owner, nil
}

// owner検索（Name）
func (uc *CarOwnerUsecase)FindByName(name string)([]*model.CarOwner, error){
    //引数のバリデーション
    if name == "" {
        return nil, fmt.Errorf("nameに検索したい名前を渡してください")
    }
    
    foundOwners, err := uc.CarOwnerRepo.FindByName(name)
    if err != nil {
        return nil, fmt.Errorf("DBのCarOwnersから対象の名前を検索するところでエラー: %w", err)
    }
    return foundOwners, nil
}



