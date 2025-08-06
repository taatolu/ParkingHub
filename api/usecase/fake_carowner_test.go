package usecase

import(
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRegistCarOwner_FakeRepo(t *testing.T){
	tests := []struct {
		testname	string
		owner		*model.CarOwner
		wantError	bool
	}{
		//testケースの作成
		{
			testname:	"正常系",
			owner:		&model.CarOwner{
				ID:1, FirstName:"test", MiddleName:"山田", LastName:"太郎",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError:	false,
		},
		{
			testname:	"異常系(FirstNameのみ存在)",
			owner:		&model.CarOwner{
				ID:1, FirstName:"test", MiddleName:"", LastName:"",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError:	true,
		},
		{
			testname:	"異常系(MiddleNameのみ存在)",
			owner:		&model.CarOwner{
				ID:1, FirstName:"", MiddleName:"山田", LastName:"",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError:	true,
		},
		{
			testname:	"異常系(LastNameのみ存在)",
			owner:		&model.CarOwner{
				ID:1, FirstName:"", MiddleName:"", LastName:"太郎",
				LicenseExpiration: time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError:	true,
		},
		{
			testname:	"異常系(免許証の期限切れ)",
			owner:		&model.CarOwner{
				ID:1, FirstName:"test", MiddleName:"山田", LastName:"太郎",
				LicenseExpiration: time.Date(2000, 11, 1, 0, 0, 0, 0, time.Local),
			},
			wantError:	true,
		},
	}
	//テストケースを回す
	for _, tt := range tests{
		t.Run(tt.testname, func(t *testing.T){
			fakeRepo := &mocks.FakeCarOwnerRepo{}
			err := fakeRepo.Save(tt.owner)
			if tt.wantError{
				//tt.wantError=true(保存できない場合)
				assert.Error(t, err, "期待していたエラーが返らない")
				assert.Nil(t, fakeRepo.SavedOwner)
			} else {
				//tt.wantError=false(保存できた場合)
				assert.NoError(t, err, "予想外のエラー")
				assert.Equal(t, tt.owner, fakeRepo.SavedOwner)
			}
		})
	}
}


//FindByIDについてはFakeテスト未実装（意味がないため）

