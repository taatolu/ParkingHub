package infrarepo

import (
	"gorm.io/gorm"
	"errors"
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
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
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


// CarOwnerを更新する
// 存在しないレコード（ID）を更新しようとするときは Error を返す
func (r *CarOwnerRepositoryImpl) Update (owner *model.CarOwner) error {
    //対象の社員がいるか確認（社員IDが一致するデータを検索）
    var count int64
    countResult := r.DB.Model(&model.CarOwner{}).Where("id = ?", owner.ID).Count(&count)
    if countResult.Error != nil {
        return fmt.Errorf("Update対象のOwnerを検索する過程でError発生: %w", countResult.Error)
    }
    //countの結果が0=>Ownerが存在しない場合
    if count == 0 {
        return fmt.Errorf("UpdateしたいOwner(ID=%v)は存在しません", owner.ID)
    }
    
    //Updateする(ID以外の項目について)
    result := r.DB.Model(&model.CarOwner{}).Where("id = ?", owner.ID).Updates(map[string]interface{}{
        "FirstName": owner.FirstName,
        "MiddleName": owner.MiddleName,
        "LastName": owner.LastName,
        "LicenseExpiration": owner.LicenseExpiration,
    })
    
    if result.Error != nil {
        return fmt.Errorf("更新失敗: %w", result.Error)
    }
    
    // 更新された行数が0の場合（条件には合致したが実際には更新されなかった場合）
    if result.RowsAffected == 0 {
        return fmt.Errorf("レコードは見つかりましたが、更新されませんでした")
    }
    
    return nil
}

// CarOwnerを削除する
// 存在しないレコード（ID）を削除しようとするときは Error を返す
func (r *CarOwnerRepositoryImpl) Delete(id uint) error {
    //引数idで渡ってくる値はuint型に変換usecase層で検証済み
    //対象のOwnerが存在するか確認
    var count int64 //int64型でないとgormのCountメソッドが受け付けない
    //carownerテーブルの中で、idが引数で渡されたidと一致するレコードの数を数える
    countResult := r.DB.Model(&model.CarOwner{}).Where("id = ?", id).Count(&count)

    //Countメソッドの実行中にエラーが発生した場合
    if countResult.Error != nil {
        return fmt.Errorf("Delete対象のOwnerを検索する過程でError発生: %w", countResult.Error)
    }

    //countの結果が0=>Ownerが存在しない場合
    if count == 0{
        return fmt.Errorf("DeleteしたいOwnerが存在しません: ID=%v", id)
    }

    //Ownerが存在する場合、削除を実行
    result := r.DB.Delete(&model.CarOwner{}, id)
    if result.Error != nil {
        return fmt.Errorf("Ownerの削除に失敗: %w", result.Error)
    }

    //削除された行数が0の場合（条件には合致したが実際には削除されなかった場合）
    if result.RowsAffected == 0 {
        return fmt.Errorf("レコードは見つかりましたが、削除されませんでした")
    }

    return nil
}
