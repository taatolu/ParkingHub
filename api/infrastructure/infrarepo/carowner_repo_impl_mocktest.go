package infrarepo

import (
	"testing"
	"fmt"
	"reflect"
	"github.com/stretchr/testify/assert"
	"time"
    "gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/taatolu/ParkingHub/api/domain/model"
)

func TestCarOwnerRepositoryImpl_Save(t *testing.T){
	//mockDBとコントローラの生成
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
        t.Fatalf("SQLiteメモリDB初期化失敗: %v", err)
    }

    // gormでテスト用テーブル作成
    db.AutoMigrate(&model.CarOwner{})

	//テスト用のモックリポジトリを設定
	repo := &CarOwnerRepositoryImpl{DB: db}
	owner := &model.CarOwner{
		ID:	1,
		FirstName:	"taro",
		MiddleName:	"山田",
		LastName:	"Yusuke",
		LicenseExpiration:	time.Now().AddDate(1, 0, 0),
	}

	//テスト対象メソッドの呼び出し
	err = repo.Save(owner)
	if err != nil {
		t.Errorf("保存に失敗: %v", err)
	}

	// DBから取得して検証
    var got model.CarOwner
    result := db.First(&got, owner.ID)
    assert.NoError(t, result.Error, "取得失敗")
    assert.Equal(t, owner.ID, got.ID, "IDが一致しない")
}


func TestCarOwnerRepositoryImpl_FindByID(t *testing.T){
    //テーブル駆動テスト(Mockテスト)
    atThisTime := time.Now()
    tests := []struct{
        testname    string
        inputID     uint
        expectError bool
        Owners []*model.CarOwner
        expectOwner *model.CarOwner
    }{
        //testケースの作成
        {
            testname:   "正常系:1件ヒット",
            inputID:    1,
            expectError:    false,
            Owners:         []*model.CarOwner{&model.CarOwner{ID:1, FirstName:"taro", MiddleName:"山田", LastName:"yusuke", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                &model.CarOwner{ID:2, FirstName:"tamaki", MiddleName:"山田", LastName:"yuichi", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            expectOwner:    &model.CarOwner{ID:1, FirstName:"taro", MiddleName:"山田", LastName:"yusuke", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
        },
        {
            testname:   "正常系:ヒット無し（IDが存在しない）",
            inputID:    3,
            expectError:    false,
             Owners:    []*model.CarOwner{&model.CarOwner{ID:1, FirstName:"taro", MiddleName:"山田", LastName:"yusuke", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                &model.CarOwner{ID:2, FirstName:"tamaki", MiddleName:"山田", LastName:"yuichi", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            expectOwner:    nil,    //ヒットしない場合はnilが返る
        },
        {
            testname:   "異常系:エラーが返る",
            inputID:    2,
            expectError:    true,
            Owners:    nil,
            expectOwner:   nil,    //ヒットしない場合はnilが返る
        },
    }
    //testケースを実行
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            //★ テストの前準備
            /// Test用にSQLiteでインメモリのDB作成
            db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
        	if err != nil {
        		t.Fatalf("sqlite初期化に失敗")
        	}
        	
        	/// テスト用のモックリポジトリを生成
        	repo := &CarOwnerRepositoryImpl{DB:	db}
            
            /// gormで接続したDB(sqliteにデータ投入)
            /// tt.Ownersが存在する場合のみCreate（存在しない状態でCreateするとエラーとなるため）
            if tt.Owners != nil { db.Create(tt.Owners) }
        	
        	//★ テスト
        	///テスト対象メソッドの呼び出し
        	got, err :=repo.FindByID(tt.inputID)
        	
        	/// エラーの有無を検証
            if tt.expectError {
                assert.Error(t, err, "エラーを期待していたが返らなかった")
            } else {
                assert.NoError(t, err, "予定外のエラーが発生しました")
            }

            /// Ownerの取得結果を検証（ポインタ型はreflect.DeepEqualが便利）
            if !reflect.DeepEqual(got, tt.expectOwner) {
                t.Errorf("取得結果が期待と異なります。got: %+v, want: %+v", got, tt.expectOwner)
            }
        })
    }
}


//DBからの絞込みができるか確認したいのでFakeテストとする
func TestCarOwnerRepositoryImpl_FindByName (t *testing.T) {
    //tableTest
    //テスト時とmock作成時に時間のずれが出ないように、現在時間を束縛
    atThisTime := time.Now()
    tests := [] struct {
        testname    string
        owners      []*model.CarOwner   //DTの中のデータ全て
        findName    string              //sqlで検索する値
        expectOwners    []*model.CarOwner       //findNameに一致するowner
        expectError bool
    }{
        //TestCaseの作成
        {
            testname:   "正常系",
            owners:     []*model.CarOwner{
                {ID:1, FirstName:"山田", MiddleName:"たかゆき", LastName:"シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID:2, FirstName:"吉松", MiddleName:"けんじ", LastName:"シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID:3, FirstName:"田中", MiddleName:"ゆうき", LastName:"ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            findName:   "田",
            expectOwners:    []*model.CarOwner{
                {ID:1, FirstName:"山田", MiddleName:"たかゆき", LastName:"シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID:3, FirstName:"田中", MiddleName:"ゆうき", LastName:"ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            expectError:    false,
        },
        {
            testname:   "正常系（検索ヒットなし）",
            owners:     []*model.CarOwner{
                {ID:1, FirstName:"山田", MiddleName:"たかゆき", LastName:"シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID:2, FirstName:"吉松", MiddleName:"けんじ", LastName:"シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID:3, FirstName:"田中", MiddleName:"ゆうき", LastName:"ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            findName:   "井",
            expectOwners:  []*model.CarOwner{},
            expectError:    false,
        },
        {
            testname:   "異常系（ownersリストがnil）",
            owners:     nil,    //DBが存在しない
            findName:   "井",
            expectOwners:   []*model.CarOwner{},
            expectError:    true,
        },
       //異常系（引数無し）をTestしようとしたが、Usecase層で引数無のバリデーションをしているので、不要とした。
    }
    //testCaseをループ処理
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            //★ テストの前準備
            /// Test用にSQLiteでインメモリのDB作成
            db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
        	if err != nil {
        		t.Fatalf("sqlite初期化に失敗")
        	}
        	
        	/// テスト用のモックリポジトリを生成
        	repo := &CarOwnerRepositoryImpl{DB:	db}
            
            /// gormで接続したDB(sqliteにデータ投入)
            /// tt.Ownersが存在する場合のみCreate（存在しない状態でCreateするとエラーとなるため）
            if tt.owners != nil { db.Create(tt.owners) }
            
            //テスト対象メソッドの呼び出し
            got, err := repo.FindByName(tt.findName)
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
            
            //Ownerの取得結果を検証
            if !reflect.DeepEqual(got, tt.expectOwners){
                t.Errorf("取得結果が期待と異なります。got: %+v, want: %+v", got, tt.expectOwners)
            }
        })
    }
}


