package usecase

import(
    "fmt"
    "github.com/taatolu/ParkingHub/api/domain/model"
    "github.com/taatolu/ParkingHub/api/domain/service"
    )

type FakeCarOwnerUsecase struct{
    //今回のFakeテストではFakeの挙動を差し替えしないので使用しないが念のため残す
    RegistCarOwnerFunc   func(owner *model.CarOwner) error
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

func (f *FakeCarOwnerUsecase) FindByID(id int) (*model.CarOwner, error) {
    //とり急ぎUsecase層の作成時にエラーが出ないよう、errorを返させる
    return nil, fmt.Errorf("FakeCarOwnerUsecase.FindByIDは未実装(handlerのテストを書くときに実装します)")
}