package postgres

import(
	"github.com/taatolu/ParkingHub/api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

//PostgresSQLへの接続初期化
func InitPostgres(conf *config.Config)(*gorm.DB, error){
	dsn := fmt.Sprintf(`host=%s dbname=%s user=%s password=%s sslmode=disable`,
		conf.DBHost, conf.DBName, conf.DBUser, conf.DBPass)
	//postgreSQLへ接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil, fmt.Errorf("postgresへの接続失敗: %w", err)
	}
	return db, nil
}
