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
            Owners:    []*model.CarOwner{&model.CarOwner{ID:1,　FirstName:"taro", MiddleName:"山田", LastName:"yusuke", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                &model.CarOwner{ID:2,　FirstName:"tamaki", MiddleName:"山田", LastName:"yuichi", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            expectOwner:    &model.CarOwner{ID:1,　FirstName:"taro", MiddleName:"山田", LastName:"yusuke", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
        },
        {
            testname:   "正常系:ヒット無し（IDが存在しない）",
            inputID:    3,
            expectError:    false,
             Owners:    []*model.CarOwner{&model.CarOwner{ID:1,　FirstName:"taro", MiddleName:"山田", LastName:"yusuke", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                &model.CarOwner{ID:2,　FirstName:"tamaki", MiddleName:"山田", LastName:"yuichi", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
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
        mockRows    *sqlmock.Rows       //findNameに一致するowner
        mockError   error
        expectError bool
    }{
        //TestCaseの作成
        {
            testname:   "正常系",
            owners:     []*model.CarOwner{
                {ID:1, FirstName:"山田", MiddleName:"たかゆき", LastName:"シニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                {ID:3, FirstName:"田中", MiddleName:"ゆうき", LastName:"ジュニア", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            findName:   "田",
            mockRows:   sqlmock.NewRows([]string{"ID", "FirstName", "MiddleName", "LastName", "LicenseExpiration"}).
                AddRow(1, "山田", "たかゆき", "シニア", atThisTime.AddDate(1, 0, 0)).
                AddRow(3, "田中", "ゆうき", "ジュニア", atThisTime.AddDate(1, 0, 0)),
            mockError:  nil,
            expectError:    false,
        },
        {
            testname:   "正常系（検索ヒットなし）",
            owners:     []*model.CarOwner{},
            findName:   "井",
            mockRows:   sqlmock.NewRows([]string{"ID", "FirstName", "MiddleName", "LastName", "LicenseExpiration"}),
            mockError:  nil,
            expectError:    false,
        },
        {
            testname:   "異常系（ownersリストがnil）",
            owners:     nil,
            findName:   "井",
            mockRows:   sqlmock.NewRows([]string{"ID", "FirstName", "MiddleName", "LastName", "LicenseExpiration"}),
            mockError:  fmt.Errorf("ownesリスト(DB)が初期化されていません"),
            expectError:    true,
        },
       //異常系（引数無し）をTestしようとしたが、Usecase層で引数無のバリデーションをしているので、不要とした。
    }
    //testCaseをループ処理
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            //sqlmockの初期化
            db, mock, err := sqlmock.New()
            if err != nil {
                t.Errorf("sqlmock初期化に失敗 %v", err)
            }
            defer db.Close()
            
            //テスト用のモックリ(sqlmock)をRepositoryImplにセットして、テスト用のリポジトリインプリメントを作成
            repo := CarOwnerRepositoryImpl{DB:db}
            
            //DBアクセスの挙動をテスト用に制御
            ///FindByName関数が実行すると予想されるSQLクエリ文（正規表現）を設定
            query := `SELECT .+ FROM carowners WHERE FirstName ILIKE \$1 OR MiddleName ILIKE \$1 OR LastName ILIKE \$1`
            
            //mockの挙動を条件分岐
            if tt.expectError {
                //異常系の場合
                mock.ExpectQuery(query).WithArgs(tt.findName).
                        WillReturnError(tt.mockError)      //queryの実行結果エラー(mockError)を返すことを期待
            } else {
                mock.ExpectQuery(query).WithArgs(tt.findName).
                        WillReturnRows(tt.mockRows)
            }
            
            //テスト対象メソッドの呼び出し
            gotOwner, err := repo.FindByName(tt.findName)
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
            
            //mockの挙動制御通りにErrorの有り無しが動くか確認
            if tt.expectError{
                //error=trueの場合
                assert.Error(t, err, "エラーが返らない")
            } else {
                assert.NoError(t, err, "予想外にエラーが返らない")
            }
            
            //Ownerの取得結果を検証
            if !reflect.DeepEqual(gotOwner, tt.owners){
                t.Errorf("取得結果が期待と異なります。got: %+v, want: %+v", gotOwner, tt.owners)
            }
            
            // SQLモックが期待を満たしているか検証
            ///mock.ExpectQuery()でセットしたものが、テスト実行中に本当に実行されたかをmock.ExpectationsWereMet()で検証
            err = mock.ExpectationsWereMet()
            if err != nil {
                t.Errorf("SQLモックの期待が満たされていません: %v", err)
            }
        })
    }
}


