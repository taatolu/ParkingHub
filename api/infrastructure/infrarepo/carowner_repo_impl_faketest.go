package infrarepo

import(
    "testing"
    "time"
    "reflect"
	"github.com/taatolu/ParkingHub/api/domain/model"
    )

// CarOwnerRepositoryImplのFakeバージョン（インメモリリストで検索）
type FakeCarOwnerRepoImpl struct {
    Owners  []*model.CarOwner   //CarOwnerのリストをDI
}

// 名前で部分一致検索するメソッド
func (fr *FakeCarOwnerRepoImpl) FindByName(name string) ([]*model.CarOwner, error) {
    var matched []*model.CarOwner
    for _, co := range fr.Owners {
        //modelに作成した、名前が含まれているか判定する関数を使用
        if co.IsContainsName(name) {
            matched = append(matched, co)
        }
    }
    return matched, nil
}

func FakeTestCarOwnerRepositoryImpl_FindByName (t *testing.T) {
    //TableTest
    //テスト時とfake作成時に時間のずれが出ないように、現在時間を束縛
    atThisTime := time.Now()
    
    tests := []struct{
        testname        string
        owners          []*model.CarOwner
        findName        string      //検索する値
        expectOwners    []*model.CarOwner
        expectError     bool
    }{
        //teatCase
        {
            testname:"正常系(2件ヒット)",
            owners:     []*model.CarOwner{
                {ID: 1, FirstName: "山田", MiddleName: "たかゆき", LastName: "シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID: 2, FirstName: "志藤", MiddleName: "ようじ", LastName: "ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID: 3, FirstName: "田中", MiddleName: "ゆうき", LastName: "ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            findName: "田",
            expectOwners:   []*model.CarOwner{
                {ID: 1, FirstName: "山田", MiddleName: "たかゆき", LastName: "シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID: 3, FirstName: "田中", MiddleName: "ゆうき", LastName: "ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            expectError: false,
        },
        {
            testname:"正常系（ヒットなし）",
            owners:     []*model.CarOwner{
                {ID: 1, FirstName: "山田", MiddleName: "たかゆき", LastName: "シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID: 2, FirstName: "志藤", MiddleName: "ようじ", LastName: "ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID: 3, FirstName: "田中", MiddleName: "ゆうき", LastName: "ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            findName: "井",
            expectOwners:   []*model.CarOwner{},
            expectError: false,
        },
    }
    //テストケースをループ処理
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            fakerepo := FakeCarOwnerRepoImpl{Owners:tt.owners}
            got, err := fakerepo.FindByName(tt.findName)
            
            if tt.expectError {
                if err == nil {
                    t.Errorf("エラーを期待していたが、エラーが返らなかった")
                }
            } else {
                if err != nil {
                    t.Errorf("予定外のエラー発生: %v", err)
                }
                // 期待IDと一致するかチェック
                var gotOwners []*model.CarOwner
                for _, o := range got {
                    gotOwners = append(gotOwners, o)
                }
                if !reflect.DeepEqual(gotOwners, tt.expectOwners) {
                    t.Errorf("結果が期待と異なる got=%v want=%v", gotOwners, tt.expectOwners)
                }
            }
        })
    }
}