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

// NameでCarOwnerを検索する
// レコードが見つからない場合は []*model.CarOwner{}, nil を返す(Errorとは返らない)
func (r *CarOwnerRepositoryImpl) FindByName (name string) ([]*model.CarOwner, error) {
    //DBかから取得した値を保存するためのリストを作成
    owners := []*model.CarOwner{}
    
    // 部分一致で検索したい場合はLIKEを使う（例えばFirstNameにnameが含まれる場合）
    // "%" + name + "%" で部分一致のパターンを作成
    // FirstNameだけでなくMiddleName, LastNameにも一致させたいなら OR 条件を使う
    err := r.DB.
        Where("first_name LIKE ? OR middle_name LIKE ? OR last_name LIKE ?", "%"+name+"%", "%"+name+"%", "%"+name+"%").
        Find(&owners).Error // エラーはここで取得
    
    if err != nil {
        return nil, fmt.Errorf("対象のOwner取得に失敗: %w", err)
    }
    
    return owners, nil
}

