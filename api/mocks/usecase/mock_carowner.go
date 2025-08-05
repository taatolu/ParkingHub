package usecase

import(
    "github.com/taatolu/ParkingHub/api/domain/model"
    )
    
//api/usecase/carowner.goのモック作成
type MockCarOwnerUsecase struct {
    //RegistCarOwnerFuncという変数をfunc(owner *model.CarOwner) error型で宣言
    //返値はError型
    RegistCarOwnerFunc func(owner *model.CarOwner) error
    FindByIDFunc func(id int)(*model.CarOwner, error)
}

func (m *MockCarOwnerUsecase) RegistCarOwner(owner *model.CarOwner) error{
    if m.RegistCarOwnerFunc != nil {
        //テストで「RegistCarOwner が呼ばれたかどうか」や
        //「呼ばれた時の動作（返すエラーなど）を自由に設定できる」
        return m.RegistCarOwnerFunc(owner)
    }
    return nil
}

func (m *MockCarOwnerUsecase)FindByID(id int) (*model.CarOwner, error) {
    return m.FindByIDFunc(id)
}