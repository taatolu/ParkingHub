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

// Saveの確認
func TestSaveCarOwner_MockRepo(t *testing.T) {
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

func TestSaveCarOwner_Error_MockRepo(t *testing.T) {
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

// FindByIDのテスト
func TestFindByID_MockRepo(t *testing.T) {
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

// FindByNameのテスト(ownerのリストを返すか？)
func TestFindByName_MockRepo(t *testing.T) {
	//tableTest
	tests := []struct{
		testname    string
		foundOwners []*model.CarOwner
		wantName    string
	}{
		{
			testname: "正常系",
			foundOwners: []*model.CarOwner{
				{ID: 1, FirstName: "test", MiddleName: "山田", LastName: "太郎",
					LicenseExpiration: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)},
				{ID: 3, FirstName: "sample", MiddleName: "横田", LastName: "健二",
					LicenseExpiration: time.Date(2025, 3, 3, 0, 0, 0, 0, time.Local)},
			},
			wantName: "田",
		},
		{
			testname: "異常系",
			foundOwners: nil,
			wantName: "",	//検索したいnameを入力していない
		},
	}
	//testCaseをループ処理
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T){
			mock := &mocks.MockCarOwnerRepo{
				FoundOwners: tt.foundOwners,
			}
			if tt.wantName == "" {
				gotOwners, err := mock.FindByName(tt.wantName)
				assert.Error(t, err, "引数が空の時、FindByNameはエラーを返すべき")
				assert.Nil(t, gotOwners, "引数が空の時gotOwnersはnilになるはず")
			} else {
				gotOwners, err := mock.FindByName(tt.wantName)
				assert.NoError(t, err, "FindByNameの正常系でエラーが発生した")
				assert.Equal(t, tt.foundOwners, gotOwners)
			}
		})
	}
}

// FindByNameシグネチャのテスト(Errorを返すか？)
func TestFindByName_Error_MockRepo(t *testing.T){
	//モックリポジトリをインスタンス化するときに、ownerのリストを渡さない(FoundOwners=nilとする)
	mock := &mocks.MockCarOwnerRepo{}
	//モックテストの本番
	gotOwners, err := mock.FindByName("test")
	assert.Error(t, err)
	assert.Nil(t, gotOwners)
}


// Updateのテスト(引数が適切に渡っているか)
func TestUpdate_MockRepo(t *testing.T) {
	//mockRepositoryの初期化
	mock := &mocks.MockCarOwnerRepo {}

	//UpdateメソッドにセットするOwnerを作成
	owner := &model.CarOwner{
		ID:	1,
		FirstName:	"やまだ",
		MiddleName:	"sample",
		LastName:	"太郎",
		LicenseExpiration:	time.Now().AddDate(1, 0, 0),
	}

	//MockRepositoryのUpdateメソッドを呼ぶ(引数にownerを渡す)
	err := mock.Update(owner)
	assert.NoError(t, err)
	assert.Equal(t, owner.ID, mock.UpdateOwner.ID)
}

// Updateのテスト(期待通りにエラーを返すか)
func TestUpdate_Error_MockRepo(t *testing.T) {
	//mockRepositoryの初期化
	mock := &mocks.MockCarOwnerRepo {}

	//MockRepositoryのUpdateメソッドを呼ぶ(引数にnil→エラーが発生)
	err := mock.Update(nil)
	assert.Error(t, err)
}
