package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/mocks"
	"testing"
	"time"
)

// /usecaseのMockテスト

// Saveシグネチャの確認
func TestSaveCarOwner(t *testing.T) {
	//mockリポジトリのインスタンス（正確には構造体のポインタ）を生成
	mockRepo := &mocks.MockCarOwnerRepo{}

	//テスト用のCarOwnerを作成
	owner := &model.CarOwner{
		ID:                1,
		FirstName:         "test",
		MiddleName:        "山田",
		LastName:          "太郎",
		LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
	}

	//mockリポジトリのSaveメソッドを呼ぶ
	err := mockRepo.Save(owner)
	if err != nil {
		//モックのインスタンスにSaveErr: errors.New("保存エラー")と化していないのでnilが返るはず
		t.Fatalf("予期しないエラー: %v", err)
	}
	//mock.Save(owner)の引数ownerがSavedOwnerに保存されたか確認
	assert.Equal(t, owner, mockRepo.SavedOwner, "SavedOwnerが正しくセットされていません")

}

func TestSaveCarOwner_Error(t *testing.T) {
	mockRepo := &mocks.MockCarOwnerRepo{
		SaveErr: errors.New("save失敗"),
	}

	owner := &model.CarOwner{
		ID: 2,
	}

	err := mockRepo.Save(owner)
	assert.Error(t, err)
	assert.Equal(t, mockRepo.SaveErr, err)
}

// FindByIDシグネチャのテスト
func TestFindByID(t *testing.T) {
	owner := &model.CarOwner{
		ID:                1,
		FirstName:         "test",
		MiddleName:        "山田",
		LastName:          "太郎",
		LicenseExpiration: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
	}

	mock := &mocks.MockCarOwnerRepo{
		FoundOwner: owner,
	}

	got, err := mock.FindByID(1)
	if err != nil {
		t.Fatalf("想定外のエラー: %v", err)
	}

	if got.ID != owner.ID ||
		got.FirstName != owner.FirstName ||
		got.MiddleName != owner.MiddleName ||
		got.LastName != owner.LastName ||
		!got.LicenseExpiration.Equal(owner.LicenseExpiration) {
		t.Errorf("取得した値が期待値と一致しません")
	}
}

// FindByNameシグネチャのテスト(ownerのリストを返すか？)
func TestFindByName(t *testing.T) {
	owners := []*model.CarOwner{
		{ID: 1, FirstName: "test", MiddleName: "山田", LastName: "太郎",
			LicenseExpiration: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)},
		{ID: 2, FirstName: "test", MiddleName: "佐藤", LastName: "一郎",
			LicenseExpiration: time.Date(2025, 2, 2, 0, 0, 0, 0, time.Local)},
		{ID: 3, FirstName: "sample", MiddleName: "横田", LastName: "健二",
			LicenseExpiration: time.Date(2025, 3, 3, 0, 0, 0, 0, time.Local)},
	}
	mock := &mocks.MockCarOwnerRepo{
		FoundOwners: owners,
	}

	//モックテストの本番処理
	//FindByName関数の引数に何を渡そうともowner=mock.FoundOwnersが帰るので、引数には適当な値（test）を渡す
	GotOwners, err := mock.FindByName("test")
	assert.NoError(t, err)
	assert.Equal(t, mock.FoundOwners, GotOwners)
}

// FindByNameシグネチャのテスト(Errorを返すか？)
func TestFindByName_Error(t *testing.T){
	//モックリポジトリをインスタンス化するときに、ownerのリストを渡さない(FoundOwners=nilとする)
	mock := &mocks.MockCarOwnerRepo{}
	//モックテストの本番
	GotOwners, err := mock.FindByName("test")
	assert.Error(t, err)
	assert.Nil(t, GotOwners)
}