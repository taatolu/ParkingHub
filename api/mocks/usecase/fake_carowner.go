package usecase

import(
    "fmt"
    "github.com/taatolu/ParkingHub/api/domain/model"
    "github.com/taatolu/ParkingHub/api/domain/service"
    "github.com/taatolu/ParkingHub/api/domain/repository"
    )

type FakeCarOwnerUsecase struct{
    //今回のFakeテストではFakeの挙動を差し替えしないので使用しないが念のため残す
    RegistCarOwnerFunc   func(owner *model.CarOwner) error
    CarOwnerRepo repository.CarOwnerRepository
}

func (f *FakeCarOwnerUsecase) RegistCarOwner(owner *model.CarOwner) error {
    //入力漏れの確認
    if !service.CarOwnerNameValidation(owner){
        //CarOwnerNameValidationがFalseだったら
		return fmt.Errorf("少なくとも２つ以上のフィールドに名前を入力ください")
    }
    
    //免許期限切れかどうかの確認
    if owner.IsLicenseExpired(){
        //IsLicenseExpiredがtrueだったら（期限切れだったら）、、、
		return fmt.Errorf("免許証期限切れの為登録不可 %v", owner.LicenseExpiration)
    }
    
    // RegistCarOwnerFuncがセットされていれば呼び出し、なければnilを返す
    if f.RegistCarOwnerFunc != nil {
        return f.RegistCarOwnerFunc(owner)
    }
    return nil
}

func (f *FakeCarOwnerUsecase) FindByID (id uint) (*model.CarOwner, error) {
    //とり急ぎUsecase層の作成時にエラーが出ないよう、errorを返させる
    return nil, fmt.Errorf("FakeCarOwnerUsecase.FindByIDは未実装(handlerのテストを書くときに実装します)")
}


func (f *FakeCarOwnerUsecase) FindByName (name string) ([]*model.CarOwner, error) {
    //引数のバリデーション
    if name == "" {
        return nil, fmt.Errorf("nameに検索したい名前を渡してください")
    }

    foundOwners, err := f.CarOwnerRepo.FindByName(name)
    if err != nil {
        return nil, fmt.Errorf("DBのCarOwnersから対象の名前を検索するところでエラー: %w", err)
    }
    return foundOwners, nil

}

func (f *FakeCarOwnerUsecase) Update (owner *model.CarOwner) error {
    //とり急ぎUsecase層の作成時にエラーが出ないよう、errorを返させる
    return fmt.Errorf("FakeCarOwnerUsecase.Updateは未実装(handlerのテストを書くときに実装します)")
}

func (f *FakeCarOwnerUsecase) Delete (id uint) error {
    //とり急ぎUsecase層の作成時にエラーが出ないよう、errorを返させる
    return fmt.Errorf("FakeCarOwnerUsecase.Deleteは未実装(handlerのテストを書くときに実装します)")
}