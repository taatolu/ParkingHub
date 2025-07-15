package mocks

import (
    "github.com/taatolu/ParkingHub/api/domain/model"
    )


//モック構造体を作成
type MockCarOwnerRepo struct{
    //リポジトリインターフェース実行の結果と照らし合わせたい内容をメンバとして定義
    SavedOwner  *model.CarOwner
    ExistUser   *model.CarOwner
    SaveErr     error
    
}


//リポジトリインターフェースのメソッドシグネチャを満たすモックのメソッドを作成
//saveメソッド
func (m *MockCarOwnerRepo) Save(carOwner *model.CarOwner) error {
    m.SavedOwner = carOwner
    return m.SaveErr
}
//deleteメソッド
func (m *MockCarOwnerRepo) Delete(id int)error{
    //処理
}

