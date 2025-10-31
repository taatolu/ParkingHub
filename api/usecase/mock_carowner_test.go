package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/mocks"
	"testing"
	"time"
)

// CarOwner構造体の比較関数
func ownerEqual(a, b *model.CarOwner) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.ID == b.ID &&
		a.FirstName == b.FirstName &&
		a.MiddleName == b.MiddleName &&
		a.LastName == b.LastName &&
		a.LicenseExpiration.Equal(b.LicenseExpiration) &&
		a.CreatedAt.Equal(b.CreatedAt) &&
		a.UpdatedAt.Equal(b.UpdatedAt)
	// 他フィールドも必要なら追加
}

// /usecaseのMockテスト

// Saveの確認
func TestSaveCarOwner_MockRepo(t *testing.T) {
	tests := []struct {
		testname  string
		owner     *model.CarOwner
		wantError bool
	}{
		//testcase
		{
			testname: "正常系",
			owner: &model.CarOwner{
				ID:                1,
				FirstName:         "test",
				MiddleName:        "山田",
				LastName:          "太郎",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError: false,
		},
		{
			testname: "異常系（名前不正）",
			owner: &model.CarOwner{
				ID:                1,
				FirstName:         "",
				MiddleName:        "",
				LastName:          "太郎",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError: true,
		},
		{
			testname: "異常系（日付不正）",
			owner: &model.CarOwner{
				ID:                1,
				FirstName:         "test",
				MiddleName:        "山田",
				LastName:          "太郎",
				LicenseExpiration: time.Now().AddDate(-1, 0, 0),
			},
			wantError: true,
		},
	}
	//testcaseをループ処理
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			//リポジトリのmockをインスタンス化
			mockRepo := &mocks.MockCarOwnerRepo{}

			//usecaseにモックのCarOwnerRepoをDI
			usecase := CarOwnerUsecase{CarOwnerRepo: mockRepo}

			//CarOwnerUsecaseのRegistCarOwnerメソッドを呼ぶ
			err := usecase.RegistCarOwner(tt.owner)

			//usecase実施後のエラーについて検証
			if tt.wantError {
				assert.Error(t, err, "なぜかエラーが発生しない")
			} else {
				assert.NoError(t, err, "予期せぬエラーが発生")

				//usecase.RegistCarOwner(owner)の引数ownerがSavedOwnerに保存されたか確認
				if !ownerEqual(tt.owner, mockRepo.SavedOwner) {
					t.Errorf("登録したOwnerと保存されたOwnerが一致しない")
				}
			}
		})
	}
}

// GetAllのテスト
func TestGetAll_MockRepo(t *testing.T){
    owners := []*model.CarOwner{
        {
            ID:                1,
    		FirstName:         "test",
    		MiddleName:        "山田",
    		LastName:          "太郎",
    		LicenseExpiration: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
        },
        {
            ID:                2,
    		FirstName:         "test",
    		MiddleName:        "山田",
    		LastName:          "はなこ",
    		LicenseExpiration: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
        },
    }
    
    mock := &mocks.MockCarOwnerRepo{
        FoundOwners:    owners,
    }
    
    //usecaseにモックのCarOwnerRepoをDI
	usecase := CarOwnerUsecase{CarOwnerRepo: mock}
    
    gotOwners, err := usecase.GetAll()
    if err != nil {
        t.Errorf("GetAllでエラー: %v", err)
    }
    assert.Equal(t, owners, gotOwners)
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
	tests := []struct {
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
			testname:    "異常系",
			foundOwners: nil,
			wantName:    "", //検索したいnameを入力していない
		},
	}
	//testCaseをループ処理
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
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
func TestFindByName_Error_MockRepo(t *testing.T) {
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
	mock := &mocks.MockCarOwnerRepo{}

	//UpdateメソッドにセットするOwnerを作成
	owner := &model.CarOwner{
		ID:                1,
		FirstName:         "やまだ",
		MiddleName:        "sample",
		LastName:          "太郎",
		LicenseExpiration: time.Now().AddDate(1, 0, 0),
	}

	//MockRepositoryのUpdateメソッドを呼ぶ(引数にownerを渡す)
	err := mock.Update(owner)
	assert.NoError(t, err)
	assert.Equal(t, owner.ID, mock.UpdateOwner.ID)
}

// Updateのテスト(期待通りにエラーを返すか)
func TestUpdate_Error_MockRepo(t *testing.T) {
	//mockRepositoryの初期化
	mock := &mocks.MockCarOwnerRepo{}

	//MockRepositoryのUpdateメソッドを呼ぶ(引数にnil→エラーが発生)
	err := mock.Update(nil)
	assert.Error(t, err)
}


// Deleteのテスト(期待通りにエラーを返すか)
func TestDelete_MockRepo(t *testing.T) {
	//テーブルテスト
	tests := []struct {
		testname  string
		id        uint
		wantError bool
	}{
		//テストケース
		{
			testname: "正常系",
			id: 1,
			wantError: false,
		},
		{
			testname: "異常系（ID未指定）",
			id: 0,
			wantError: true,
		},
	}
	//testcaseをループ処理
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T){
			//mockRepositoryの初期化
			mock := &mocks.MockCarOwnerRepo{}

			//usecase層のDeleteメソッドを呼ぶ
			usecase := CarOwnerUsecase{CarOwnerRepo: mock}
			err := usecase.Delete(tt.id)

			//errorの確認
			if tt.wantError {
				assert.Error(t, err, "なぜかエラーが発生しない")
			} else {
				assert.NoError(t, err, "予期せぬエラーが発生")
				assert.Equal(t, tt.id, mock.DeleteID, "Deleteメソッドに渡したIDがセットされていない")
			}
		})
	}
}