package infrarepo

import (
	"testing"
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock作成に失敗")
	}
	defer db.Close()

	//テスト用のモックリポジトリを生成
	repo := &CarOwnerRepositoryImpl{
		DB:	db,
	}
	targetOwner:=&model.CarOwner{
		ID:	1,
		FirstName:	"taro",
		MiddleName:	"山田",
		LastName:	"Yusuke",
		LicenseExpiration:	time.Now().AddDate(1, 0, 0),
	}

	//期待するSQLと返り値の設定
	///期待する返り値
	rows := sqlmock.NewRows([]string{"ID", "FirstName", "MiddleName", "LastName", "LicenseExpiration"}).
		AddRow(targetOwner.ID, targetOwner.FirstName, targetOwner.MiddleName, targetOwner.LastName, targetOwner.LicenseExpiration)

	mock.ExpectQuery("SELECT (.+) FROM carowners WHERE ID = \\$1").
		WithArgs(targetOwner.ID).
		WillReturnRows(rows)
	
	//テスト対象メソッドの呼び出し
	gotOwner, err := repo.FindByID(targetOwner.ID)
	if err != nil {
		t.Errorf("取得失敗: %v", err)
	}
	assert.Equal(t, targetOwner.ID, gotOwner.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("期待したSQLが実行されていません: %v", err)
	}
}