package usecase

import (
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/domain/repository"
	"github.com/taatolu/ParkingHub/api/domain/service"
	"gorm.io/gorm"
)

// ユースケース構造体の作成
type CarOwnerUsecase struct {
	CarOwnerRepo repository.CarOwnerRepository
}

// CarOwnerUsecaseのインターフェースを作成(後に作成するHandlerのテストの際にUsecseの際し変えが見込まれるから)
type CarOwnerUsecaseIF interface {
	RegistCarOwner(owner *model.CarOwner) error
	GetAll() ([]*model.CarOwner, error)
	FindByID(id uint) (*model.CarOwner, error)
	FindByName(name string) ([]*model.CarOwner, error)
	Update(owner *model.CarOwner) error
	Delete(id uint) error
}

// owner登録処理
func (uc *CarOwnerUsecase) RegistCarOwner(owner *model.CarOwner) error {
	//入力漏れの確認
	if !service.CarOwnerNameValidation(owner) {
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

// owner全件取得
func (uc *CarOwnerUsecase) GetAll() ([]*model.CarOwner, error) {
    // インフラストラクチャ層のGetAllを実行
    owners, err := uc.CarOwnerRepo.GetAll()
	if err != nil{
		return nil, fmt.Errorf("DBのCarOwnersから全件取得するところでエラー: %w", err)
	}

	// ownersが空の場合の処理
	// 空リストは正常なレスポンスとして返す
	return owners, nil
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
	if err == gorm.ErrRecordNotFound {
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
func (uc *CarOwnerUsecase) FindByName(name string) ([]*model.CarOwner, error) {
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

// ownerのUpdate
func (uc *CarOwnerUsecase) Update(carOwner *model.CarOwner) error {
	//引数が渡されていない場合
	if carOwner == nil {
		return fmt.Errorf("carOwnerが引数で渡されていません")
	}

	//引数で渡されたCarOwnerのIDが整数か検証
	if !carOwner.IsIDPositive() {
		return fmt.Errorf("引数で渡されたCarOwnerのIDが負の数です")
	}

	//更新しようとするOwnerのname項目に不足がないか
	if !service.CarOwnerNameValidation(carOwner) {
		//CarOwnerNameValidationがFalseだったら
		return fmt.Errorf("少なくとも２つ以上のフィールドに名前を入力ください")
	}

	//値が更新されていない場合
	existingOwner, err := uc.FindByID(carOwner.ID)
	if err != nil {
		return fmt.Errorf("現在のOwnerの取得に失敗: %w", err)
	}
	if existingOwner.ID == carOwner.ID &&
		existingOwner.FirstName == carOwner.FirstName &&
		existingOwner.MiddleName == carOwner.MiddleName &&
		existingOwner.LastName == carOwner.LastName &&
		existingOwner.LicenseExpiration == carOwner.LicenseExpiration {
		return fmt.Errorf("現在のCarOwnerと変更したい値に差がありません")
	}

	//免許証の期限更新をしようとする時に期限切れがないか確認
	//元々登録されていた免許期限が切れていて、今回の更新が免許更新でない場合はエラーを吐かない
	if existingOwner.LicenseExpiration != carOwner.LicenseExpiration {
		//元々登録されていた免許期限 != 新規に登録使用する免許期限(免許期限を更新しようとする時)
		if carOwner.IsLicenseExpired() {
			//新規に登録しようとする免許期限日が、すでに免許切れの場合
			return fmt.Errorf("更新しようとする免許期限の日付が免許切れです: %v", carOwner.LicenseExpiration)
		}
	}

	//domainRepositoryを使って更新
	// エラーが返る場合:
	// - DB接続エラー
	// - 指定したIDのCarOwnerが存在しない場合
	// - 更新処理中に予期せぬ例外が発生した場合
	return uc.CarOwnerRepo.Update(carOwner)
}

// Delete(Owner削除)
func (uc *CarOwnerUsecase) Delete(id uint) error {
	//引数のバリデーション
	if id <= 0 {
		return fmt.Errorf("削除したい値は０より大きい整数をセットしてください")
	}
	//repository層のDeleteメソッドを呼ぶ
	return uc.CarOwnerRepo.Delete(id)
}