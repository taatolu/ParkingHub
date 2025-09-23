package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/mocks"
	"testing"
	"time"
)

func TestRegistCarOwner_FakeRepo(t *testing.T) {
	tests := []struct {
		testname  string
		owner     *model.CarOwner
		wantError bool
	}{
		//testケースの作成
		{
			testname: "正常系",
			owner: &model.CarOwner{
				ID: 1, FirstName: "test", MiddleName: "山田", LastName: "太郎",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError: false,
		},
		{
			testname: "異常系(FirstNameのみ存在)",
			owner: &model.CarOwner{
				ID: 1, FirstName: "test", MiddleName: "", LastName: "",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError: true,
		},
		{
			testname: "異常系(MiddleNameのみ存在)",
			owner: &model.CarOwner{
				ID: 1, FirstName: "", MiddleName: "山田", LastName: "",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError: true,
		},
		{
			testname: "異常系(LastNameのみ存在)",
			owner: &model.CarOwner{
				ID: 1, FirstName: "", MiddleName: "", LastName: "太郎",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError: true,
		},
		{
			testname: "異常系(免許証の期限切れ)",
			owner: &model.CarOwner{
				ID: 1, FirstName: "test", MiddleName: "山田", LastName: "太郎",
				LicenseExpiration: time.Date(2000, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError: true,
		},
	}
	//テストケースを回す
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			fakeRepo := &mocks.FakeCarOwnerRepo{}
			usecase := CarOwnerUsecase{CarOwnerRepo: fakeRepo}
			err := usecase.RegistCarOwner(tt.owner)
			if tt.wantError {
				//tt.wantError=true(保存できない場合)
				assert.Error(t, err, "期待していたエラーが返らない")
			} else {
				//tt.wantError=false(保存できた場合)
				assert.NoError(t, err, "予想外のエラー")
				assert.Equal(t, tt.owner, fakeRepo.SavedOwner)
			}
		})
	}
}

// FindByIDのテスト（引数の型チェックのみ実装）
func TestFindByID_FakeRepo(t *testing.T) {
	tests := []struct {
		testname  string
		id        uint
		wantError bool
	}{
		//テストケースの作成
		{
			testname:  "正常系",
			id:        1,
			wantError: false,
		},
		{
			testname:  "異常系",
			id:        0, //idはuint型なので負の数にはなりえない
			wantError: true,
		},
	}
	//テスト実行
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			fakeRepo := &mocks.FakeCarOwnerRepo{}
			carOwner, err := fakeRepo.FindByID(tt.id)
			if tt.wantError {
				//異常系の場合
				assert.Error(t, err, "IDが0なのでエラーを期待したが、エラーにならない")
			} else {
				assert.NoError(t, err, "エラーが発生してしまった")
				assert.NotNil(t, carOwner, "carOwnerがnilです")
				assert.Equal(t, tt.id, carOwner.ID)
			}
		})
	}
}

//FindByNameのFakeTestは実装しない（必要に応じて追加）
///ハンドラから渡されるパラメータのテスト
///-引数が空であれば、パラメータ不正としてハンドラ側でエラーとなり、Usecaseまで渡ってこない

// Updateのテスト
func TestUpdate_FakeRepo(t *testing.T) {
	//tableTest
	tests := []struct {
		testname          string
		existingOwner     *model.CarOwner //元々登録されていたOwner
		updateOwner       *model.CarOwner //更新したいOwner
		repoImplFindError bool            //repositoryImplのFindByIDメソッドから返されるError
		wantError         bool            //UsecaseのUpdate実行結果返るエラー
	}{
		//testCaseの作成
		{
			testname:          "正常系",
			existingOwner:     &model.CarOwner{ID: 1, FirstName: "山田", MiddleName: "太郎", LastName: "junior", LicenseExpiration: time.Now().AddDate(1, 0, 0)},
			updateOwner:       &model.CarOwner{ID: 1, FirstName: "花田", MiddleName: "太郎", LastName: "senior", LicenseExpiration: time.Now().AddDate(2, 0, 0)},
			repoImplFindError: false, //repositoryImplのFindByIDメソッドからErrorは帰らない
			wantError:         false,
		},
		{
			testname:          "異常系(存在しないIDの更新)",
			existingOwner:     &model.CarOwner{ID: 1, FirstName: "山田", MiddleName: "太郎", LastName: "junior", LicenseExpiration: time.Now().AddDate(1, 0, 0)},
			updateOwner:       &model.CarOwner{ID: 2, FirstName: "花田", MiddleName: "太郎", LastName: "senior", LicenseExpiration: time.Now().AddDate(2, 0, 0)},
			repoImplFindError: true, //repositoryImplのFindByIDメソッドからgorm.ErrRecordNotFoundが返される
			wantError:         true,
		},
		{
			testname:          "異常系(更新するname項目の不足)",
			existingOwner:     &model.CarOwner{ID: 1, FirstName: "山田", MiddleName: "太郎", LastName: "junior", LicenseExpiration: time.Now().AddDate(1, 0, 0)},
			updateOwner:       &model.CarOwner{ID: 1, FirstName: "花田", MiddleName: "", LastName: "", LicenseExpiration: time.Now().AddDate(2, 0, 0)},
			repoImplFindError: false, //repositoryImplのFindByIDメソッドからエラーは帰らない
			wantError:         true,
		},
		{
			testname:          "異常系(更新する免許期限の期限切れ)",
			existingOwner:     &model.CarOwner{ID: 1, FirstName: "山田", MiddleName: "太郎", LastName: "junior", LicenseExpiration: time.Now().AddDate(1, 0, 0)},
			updateOwner:       &model.CarOwner{ID: 1, FirstName: "山田", MiddleName: "太郎", LastName: "junior", LicenseExpiration: time.Now().AddDate(-1, 0, 0)},
			repoImplFindError: false,
			wantError:         true,
		},
	}
	//testCaseをループ処理
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			//FakeのRepositoryImplを実装
			//Usecaseの（Updae）Fakeテストの中で、FindByIDを実行するので、
			//あたかもInfrarepoから返されたように
			repo := &mocks.FakeCarOwnerRepo{SavedOwner: tt.existingOwner, WantError: tt.repoImplFindError}

			//本番のUsecaseにFakeのRepositoryImplをDI
			usecase := CarOwnerUsecase{CarOwnerRepo: repo}

			//本番のUsecaseの挙動を確認
			err := usecase.Update(tt.updateOwner)

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
