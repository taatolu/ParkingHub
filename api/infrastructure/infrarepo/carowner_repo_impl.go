package infrarepo

import (
	"gorm.io/gorm"
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
)

type CarOwnerRepositoryImpl struct{
	DB *gorm.DB
}

//repositoryの具体的な実装部分
///CarOwner_Save
func (r *CarOwnerRepositoryImpl) Save (carOwner *model.CarOwner) error {

	err := r.DB.Create(carOwner).Error
	if err != nil{
		return fmt.Errorf("CarOwnerのデータ登録失敗: %w", err)
	}
	
	return nil
}

// IDでCarOwnerを検索する
// レコードが見つからない場合は nil, gorm.ErrRecordNotFound を返す
func (r *CarOwnerRepositoryImpl) FindByID(id uint) (*model.CarOwner, error) {
    // CarOwnerモデルの値（検索結果を格納する箱）を用意
    var owner model.CarOwner

    // 指定IDでCarOwnerを1件検索
    result := r.DB.First(&owner, id)

    // レコードが見つからなかった場合
    if result.Error == gorm.ErrRecordNotFound {
        // データがなかった場合は nil, gorm.ErrRecordNotFound を返す
        return nil, gorm.ErrRecordNotFound
    }

    // その他のエラー（DB接続失敗など）があれば、そのまま返す
    if result.Error != nil {
        return nil, fmt.Errorf("CarOwnerの取得失敗: %w", result.Error)
    }

    // 正常に取得できた場合は、ownerのポインタとnil（エラーなし）を返す
    return &owner, nil
}


func (r *CarOwnerRepositoryImpl) FindByName (name string) ([]*model.CarOwner, error) {
    //DBかから取得した値を保存するためのリストを作成
    owners := []*model.CarOwner{}
    
    cmd := `SELECT ID, FirstName, MiddleName, LastName, LicenseExpiration FROM carowners WHERE FirstName ILIKE $1 OR MiddleName ILIKE $1 OR LastName ILIKE $1`
    
    //DBから一致するデータを(rowsに)取得
    rows, err := r.DB.Query(cmd, name)
    if err != nil {
        return nil, fmt.Errorf("DBからの取得に失敗: %w", err)
    }
    defer rows.Close()
    
    for rows.Next() {
        //取得したrowsからownerを一つ取り出して保存するための変数作成
        owner := &model.CarOwner{}
        //上記変数に保存
        err = rows.Scan(
            &owner.ID,
            &owner.FirstName,
            &owner.MiddleName,
            &owner.LastName,
            &owner.LicenseExpiration,)
        if err != nil {
            return nil, fmt.Errorf("対象のOwner取得に失敗: %w", err)
        }
        //errorなくここまで進んだら、作成しておいたowners配列にOwnerを保存
        owners = append(owners, owner)
    }
    return owners, nil
}

