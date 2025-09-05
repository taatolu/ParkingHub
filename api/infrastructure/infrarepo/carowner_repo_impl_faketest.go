package infrarepo

import (
    "testing"
    "github.com/taatolu/ParkingHub/api/domain/model"
    mockRepo "github.com/taatolu/ParkingHub/api/mocks/repository"
    "time"
    "reflect"
)

// CarOwnerRepositoryImplのFakeを使用したFindByNameのテスト)(インメモリリストで検索)
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
        //testCase
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
            fakerepo := mockRepo.FakeCarOwnerRepoImpl{Owners:tt.owners}        //ここでFakeを作成
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
                if !reflect.DeepEqual(got, tt.expectOwners) {
                    t.Errorf("結果が期待と異なる got=%v want=%v", got, tt.expectOwners)
                }
            }
        })
    }
}