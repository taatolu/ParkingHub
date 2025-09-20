package infrarepo

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
    "gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/taatolu/ParkingHub/api/domain/model"
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
            Owners:         []*model.CarOwner{
                &model.CarOwner{ID:1, FirstName:"taro", MiddleName:"山田", LastName:"yusuke", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
                &model.CarOwner{ID:2, FirstName:"tamaki", MiddleName:"山田", LastName:"yuichi", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
            },
            expectOwner:    &model.CarOwner{ID:1, FirstName:"taro", MiddleName:"山田", LastName:"yusuke", LicenseExpiration: atThisTime.AddDate(1, 0, 0)},
        },
        {
            testname:   "正常系:ヒット無し（IDが存在しない）",
            inputID:    3,
            expectError:    true,       //RecordNotFoundのエラーが返る使用だから
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
        tt := tt // クロージャキャプチャ対策として、ループ変数を関数スコープにコピー
        t.Run(tt.testname, func(t *testing.T){
            //★ テストの前準備
            /// Test用にSQLiteでインメモリのDB作成
            db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
        	if err != nil {
        		t.Fatalf("sqlite初期化に失敗")
        	}
        	
        	// gormでテスト用テーブル作成
            db.AutoMigrate(&model.CarOwner{})
        	
        	/// テスト用のモックリポジトリを生成
        	repo := &CarOwnerRepositoryImpl{DB:	db}
            
            /// gormで接続したDB(sqliteにデータ投入)
            /// tt.Ownersが存在する場合のみCreate（存在しない状態でCreateするとエラーとなるため）
            if tt.Owners != nil {
                for _, o := range tt.Owners {
                    db.Create(o) // 1件ずつ登録
                }
            }
        	
        	//★ テスト
        	///テスト対象メソッドの呼び出し
        	got, err :=repo.FindByID(tt.inputID)
        	
        	/// エラーの有無を検証
            if tt.expectError {
                assert.Error(t, err, "エラーを期待していたが返らなかった")
            } else {
                assert.NoError(t, err, "予定外のエラーが発生しました")
            }
            
            //gotがnilでなかったら
            if got != nil {
                // DB登録後の値で期待値をセット
                // gormが自動でCreatedAt等の値を作成するため、実際にDBから取得した値で期待値をセットする
                tt.expectOwner.CreatedAt = got.CreatedAt
                tt.expectOwner.UpdatedAt = got.UpdatedAt
                
                /// Ownerの取得結果を検証
                if !ownerEqual(got, tt.expectOwner) {
                        t.Errorf("取得結果が期待と異なります。\ngot: %+v\nwant: %+v", got, tt.expectOwner)
                    }
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
            //このowners=nil つまりDBが存在しない場合は、本番実装のr.DB.Where～を通ってもerrには何も返らず、&ownersに空スライスが返る
            findName:   "井",
            expectOwners:   []*model.CarOwner{},
            expectError:    true,
        },
       //異常系（引数無し）をTestしようとしたが、Usecase層で引数無のバリデーションをしているので、不要とした。
    }
    //testCaseをループ処理
    for _, tt := range tests {
        tt := tt // クロージャキャプチャ対策として、ループ変数を関数スコープにコピー
        t.Run(tt.testname, func(t *testing.T){
            //★ テストの前準備
            /// Test用にSQLiteでインメモリのDB作成
            db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
        	if err != nil {
        		t.Fatalf("sqlite初期化に失敗")
        	}
        	
        	// tt.Ownersが存在する場合のみgormでテスト用テーブル作成
            if tt.owners != nil {db.AutoMigrate(&model.CarOwner{})}
        	
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
            
            // 期待値リストと実際の取得結果リストを比較する例
            for i, gotOwner := range got {
                // 期待値リスト expectOwners の i 番目を取得
                expectOwner := tt.expectOwners[i]
            
                // フィールドを比較
                // ここでは比較例
                if gotOwner.ID != expectOwner.ID {
                    t.Errorf("取得したOwnerのIDが一致しません: got=%v, expect=%v", gotOwner.ID, expectOwner.ID)
                }
            }
            
            // 最後にdb削除（他のテストに影響しないように）
            err = db.Migrator().DropTable("car_owners")
        })
    }
}


// TestCarOwnerRepositoryImpl_Update は CarOwnerRepositoryImpl の Update メソッドが
// 正しく既存のオーナー情報を更新できるか、また存在しないオーナーを更新しようとした場合に
// エラーが返るかを検証するテストです。
func TestCarOwnerRepositoryImpl_Update (t *testing.T) {
    //tableTestの作成
    tests := []struct {
        testname        string
        existingOwner   *model.CarOwner     //DBに初期に設置されるOwner
        updateOwner     *model.CarOwner     //変更される値
        wantError       bool
    }{
        //testCaseの作成
        {
            testname:       "正常系",
            existingOwner:  &model.CarOwner{ID:1, FirstName:"yamada", MiddleName:"takayuki", LastName:"junior"},
            updateOwner:    &model.CarOwner{ID:1, FirstName:"yamada", MiddleName:"takayuki", LastName:"senior"},
            wantError:      false,
        },
        {
            testname:       "異常系（存在しないOwner(ID:2)を更新しようとする）",
            existingOwner:  &model.CarOwner{ID:1, FirstName:"yamada", MiddleName:"takayuki", LastName:"junior"},
            updateOwner:    &model.CarOwner{ID:2, FirstName:"yamada", MiddleName:"takayuki", LastName:"senior"},
            wantError:      true,
        },
    }
    //testCaseをループ処理
    for _, tt := range tests {
        tt := tt    // クロージャキャプチャ対策として、ループ変数を関数スコープにコピー
        t.Run(tt.testname, func(t *testing.T){
            //testの前準備（DB準備）sqliteのDBをインメモリで作成
            db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
            if err != nil{
                t.Fatalf("sqliteの初期化に失敗")
            }
            
            //gormでテスト用テーブル作成
            if err := db.AutoMigrate(&model.CarOwner{}); err != nil{
                t.Fatalf("テスト用のテーブルの作成失敗")
            }

            //上記テーブルにデータ投入
            if err := db.Create(tt.existingOwner).Error; err != nil {
                t.Fatalf("テスト用データの初期設定に失敗")
            }

            //本番で作成したリポジトリの実装（RepositoryImpl）にインメモリのDB(sqlite)をセット
            repo := &CarOwnerRepositoryImpl{DB: db}

            //テスト対象メソッドの実行
            updateErr := repo.Update(tt.updateOwner)
            //実行結果のエラー判定
            if tt.wantError {
                assert.Error(t, updateErr)
            } else {
                assert.NoError(t, updateErr)
            }
            
            //保存したデータがUpsertされているか確認
            ///更新後のデータを取得
            got, findErr := repo.FindByID(tt.updateOwner.ID)
            
            if tt.wantError {
                //wantErrorがtrueのとき
                assert.Error(t,findErr)
            } else {
                //wantErrorがfalseのとき
                assert.NoError(t, findErr)
                
                //変更後用として渡したデータが変更後に取得したデータと一致するか確認
                ///gormが自動でセットする項目は引き渡し
                tt.updateOwner.CreatedAt = got.CreatedAt
                tt.updateOwner.UpdatedAt = got.UpdatedAt
                
                ///確認用の関数で確認
                if !ownerEqual(got, tt.updateOwner) {
                    t.Errorf("取得結果が期待と異なります。\ngot: %+v\nwant: %+v", got, tt.updateOwner)
                }
            }
        })
    }
}

