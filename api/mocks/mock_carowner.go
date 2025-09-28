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
	UpdateOwner *model.CarOwner
	DeleteID    uint
}

// リポジトリインターフェースのメソッドシグネチャを満たすモックのメソッドを作成
// saveメソッド
func (m *MockCarOwnerRepo) Save(carOwner *model.CarOwner) error {
	if carOwner != nil {
		m.SavedOwner = carOwner
		return nil
	}
	return fmt.Errorf("保存失敗")
}

// FindByID
func (m *MockCarOwnerRepo) FindByID(id uint) (*model.CarOwner, error) {
	return m.FoundOwner, nil
}

// FindByName(部分一致検索)
func (m *MockCarOwnerRepo) FindByName(name string) ([]*model.CarOwner, error) {
	if m.FoundOwners == nil{
        return nil, fmt.Errorf("検索値にヒットするOwnrは存在しません")
    }
    return m.FoundOwners, nil
}


// Update
func (m *MockCarOwnerRepo) Update (carOwner *model.CarOwner) error {
	m.UpdateOwner = carOwner
	if m.UpdateOwner == nil {
		return fmt.Errorf("更新したい値をセットしてください")
	}
	return nil
}

//Delete
func (m *MockCarOwnerRepo) Delete(id uint) error {
	//usecase層でIDが整数であるようにバリデーションしているので、repository層ではIDが0でないことを確認するだけで良い
	//本番のrepository層でも同様の判定を何処かに入れた方が良いかも
	if id == 0 {
		return fmt.Errorf("削除したいIDをセットしてください")
	}

	m.DeleteID = id
	return nil
}