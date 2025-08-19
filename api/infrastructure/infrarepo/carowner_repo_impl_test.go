package infrarepo

import (
	"testing"
	"fmt"
	"reflect"
	"github.com/stretchr/testify/assert"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/taatolu/ParkingHub/api/domain/model"
)

func TestCarOwnerRepositoryImpl_Save(t *testing.T){
	//mockDBとコントローラの生成
	db, mock, err := sqlmock.New()
	if err != nil{
		t.Fatalf("sqlmock作成失敗")
	}
	defer db.Close()

	//テスト用のモックリポジトリを設定
	repo := &CarOwnerRepositoryImpl{DB: db}
	owner := &model.CarOwner{
		ID:	1,
		FirstName:	"taro",
		MiddleName:	"山田",
		LastName:	"Yusuke",
		LicenseExpiration:	time.Now().AddDate(1, 0, 0),
	}

	//期待するSQL・引数・返り値の設定
	mock.ExpectExec("INSERT INTO carowners").
	WithArgs(owner.ID, owner.FirstName, owner.MiddleName, owner.LastName, owner.LicenseExpiration).
	WillReturnResult(sqlmock.NewResult(1, 1))

	//テスト対象メソッドの呼び出し
	err = repo.Save(owner)
	if err != nil {
		t.Errorf("保存に失敗: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("期待したSQLが実行されていません: %v", err)
	}
}


func TestCarOwnerRepositoryImpl_FindByID(t *testing.T){
    //テーブル駆動テスト(Mockテスト)
    tests := []struct{
        testname    string
        inputID     int
        mockRows     *sqlmock.Rows             // モックで返す行データ
        mockError    error                     // モックで返すエラ
        expectError bool
        expectOwner *model.CarOwner
    }{
        //testケースの作成
        {
            testname:   "正常系:1件ヒット",
            inputID:    1,
            mockRows:   sqlmock.NewRows([]string{"ID", "FirstName", "MiddleName", "LastName", "LicenseExpiration"}).
                AddRow(1, "taro", "山田", "yusuke", time.Now().AddDate(1, 0, 0)),
            expectError:    false,
            expectOwner:    &model.CarOwner{
                ID: 1,
                FirstName:  "taro",
        	    MiddleName: "山田",
        	    LastName:   "yusuke",
        	    LicenseExpiration: time.Now().AddDate(1, 0, 0),
            },
        },
        {
            testname:   "正常系:ヒット無し（IDが存在しない）",
            inputID:    2,
            mockRows:   sqlmock.NewRows([]string{"ID", "FirstName", "MiddleName", "LastName", "LicenseExpiration"}),
            //レコード無し
            expectError:    false,
            expectOwner:    nil,
        },
        {
            testname:   "異常系:エラーが返る",
            inputID:    3,
            mockRows:   nil,    // エラーなのでrowsを返さずnilが返る
            mockError: fmt.Errorf("DB接続失敗"),
            expectError:    true,
            expectOwner:    nil,
        },
    }
    //testケースを実行
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            db, mock, err := sqlmock.New()
        	if err != nil {
        		t.Fatalf("sqlmock作成に失敗")
        	}
        	defer db.Close()
        	
        	//テスト用のモックリポジトリを生成
        	repo := &CarOwnerRepositoryImpl{DB:	db}
        	
        	//SQLモックのセット
        	query := "SELECT (.+) from carowners WHERE id = \\$1"
        	
        	if tt.mockError != nil {
        	    //クエリの実行結果Errorを返すように作成
        	    mock.ExpectQuery(query).
        	        WithArgs(tt.inputID).
        	        WillReturnError(tt.mockError)
        	} else {
        	    //クエリの実行結果Rowを返すように作成
        	    mock.ExpectQuery(query).
        	        WithArgs(tt.inputID).
        	        WillReturnRows(tt.mockRows)
        	}
        	
        	//テスト対象メソッドの呼び出し
        	gotOwner, err :=repo.FindByID(tt.inputID)
        	
        	//errorが発生するかどうかの確認
        	if tt.expectError{
        	    assert.Error(t, err, "エラーを期待していたがエラーが返らない")
        	} else {
        	    assert.NoError(t, err , "予定外にエラーが発生しました")
        	}
        	
        	// Ownerの取得結果を検証
            if !reflect.DeepEqual(gotOwner, tt.expectOwner) {
                t.Errorf("取得結果が期待と異なります。got: %+v, want: %+v", gotOwner, tt.expectOwner)
            }

            // SQLモックの期待を満たしているか検証
            ///mock.ExpectQuery()でセットしたものが、テスト実行中に本当に実行されたかをmock.ExpectationsWereMet()で検証
            if err := mock.ExpectationsWereMet(); err != nil {
                t.Errorf("SQLモックの期待が満たされていません: %v", err)
            }
        })
    }

}