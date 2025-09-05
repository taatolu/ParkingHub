package repository

import (
    "github.com/taatolu/ParkingHub/api/domain/model"
)


// CarOwnerRepositoryImplのFakeバージョン（インメモリリストで検索）
type FakeCarOwnerRepoImpl struct {
    Owners  []*model.CarOwner   //CarOwnerのリストをDI
}

// FindByName は名前で部分一致検索するメソッドです。
// 引数: name - 検索したい名前の文字列（部分一致）
// 返り値: 一致したCarOwnerのスライスとエラー（通常はnil）
func (fr *FakeCarOwnerRepoImpl) FindByName(name string) ([]*model.CarOwner, error) {
    var matched []*model.CarOwner
    for _, co := range fr.Owners {
        // modelに作成した、名前が含まれているか判定する関数を使用
        if co.IsContainsName(name) {
            matched = append(matched, co)
        }
    }
    return matched, nil
}