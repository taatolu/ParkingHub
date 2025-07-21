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