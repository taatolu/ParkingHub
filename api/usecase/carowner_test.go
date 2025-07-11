package usecase

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
    "api/mocks"
    "api/model"
    )

//Saveシグネチャの確認
func TestRegisterUser(t *testing.T){
    //処理
}

func TestSaveCarQoner(t *testing.T){
    //mockリポジトリのインスタンス（正確には構造体のポインタ）を生成
    mocRepo := &mocks.MockCarOwnerRepo{}
    
    //テスト用のCarOwnerを作成
    owner := &model.CarOwner{
        ID: 1,
        FirstName:  "test",
        MiddleName: "山田",
        LastName:   "太郎",
        LicenseExpiration:  time.Date(2025, 11, 1, 0, 0, 0, 0, time.Local),
    }
    
    //mockリポジトリのSaveメソッドを呼ぶ
    err := mocRepo.Save(owner)
    if err != nil{
        //モックのインスタンスにSaveErr: errors.New("保存エラー")と化していないのでnilが返るはず
        t.Fatalf("予期しないエラー: %v", err)
    }
    //mock.Save(owner)の引数ownerがSavedOwnerに保存されたか確認
    assert.Equal(t, owner, mocRepo.SavedOwner, "SavedOwnerが正しくセットされていません")
    
}