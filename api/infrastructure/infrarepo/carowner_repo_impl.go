package infrarepo

import (
	"database/sql"
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
)

type CarOwnerRepositoryImpl struct{
	DB *sql.DB
}

//repositoryの具体的な実装部分
///CarOwner_Save
func (r *CarOwnerRepositoryImpl) Save (carOwner *model.CarOwner) error {
	cmd := `INSERT INTO carowners (
	ID,
	FirstName,
	MiddleName,
	LastName,
	LicenseExpiration) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.DB.Exec(cmd,
		carOwner.ID,
		carOwner.FirstName,
		carOwner.MiddleName,
		carOwner.LastName,
		carOwner.LicenseExpiration)
	if err != nil{
		return fmt.Errorf("CarOwnerのデータ登録失敗: %w", err)
	}
	
	return nil
}

func (r *CarOwnerRepositoryImpl) FindByID (id int)(*model.CarOwner, error){
	//DBから取得したownerを格納する場所を作成
	owner := &model.CarOwner{}
	
	cmd := `SELECT ID, FirstName, MiddleName, LastName, LicenseExpiration FROM carowners WHERE id = $1`
	err := r.DB.QueryRow(cmd, id).Scan(
	    &owner.ID,
	    &owner.FirstName,
	    &owner.MiddleName,
	    &owner.LastName,
	    &owner.LicenseExpiration)
	
	//sql実行の結果、該当するデータが無ければsql: no rows in result setとエラーが帰ってしまう
	//本質的にはerrorではないので、model.CarOwner=nil, error=nilとして返したい
	if err == sql.ErrNoRows {
	    return nil, nil     //データが存在しない場合はnilを返す
	}
	
	if err != nil{
	    return nil, fmt.Errorf("CarOwnerの取得失敗: %w", err)
	}
	return owner, nil
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

