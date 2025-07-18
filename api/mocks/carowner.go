package mocks

import (
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
)

// モック構造体を作成
type MockCarOwnerRepo struct {
	//リポジトリインターフェース実行の結果と照らし合わせたい内容をメンバとして定義
	SavedOwner *model.CarOwner
	SaveErr    error
	FoundOwner *model.CarOwner
    FoundOwners []*model.CarOwner
}

// リポジトリインターフェースのメソッドシグネチャを満たすモックのメソッドを作成
// saveメソッド
func (m *MockCarOwnerRepo) Save(carOwner *model.CarOwner) error {
	m.SavedOwner = carOwner
	return m.SaveErr
}

// FindByID
func (m *MockCarOwnerRepo) FindByID(id int) (*model.CarOwner, error) {
	return m.FoundOwner, nil
}

// FindByName(部分一致検索)
func (m *MockCarOwnerRepo) FindByName(name string) ([]*model.CarOwner, error) {
	if m.FoundOwners == nil{
        return nil, fmt.Errorf("検索値にヒットするOwnrは存在しません")
    }
    return m.FoundOwners, nil
}
