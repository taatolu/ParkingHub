package migrate

import(
	"log"
	"fmt"
	"github.com/taatolu/ParkingHub/api/config"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/infrastructure/postgres"
)

func RunMigration(conf *config.Config) error {
    db, err := postgres.InitPostgres(conf)
    if err != nil {
        return fmt.Errorf("dbの初期化失敗: %w", err)
    }
    //InitPostgresから受け取ったdb(gorm.DB)でAutoMigrate
	// AutoMigrateでテーブル作成・更新
	// ここで必要なテーブル構造体を渡します
	if err := db.AutoMigrate(&model.CarOwner{}); err != nil {
		return fmt.Errorf("マイグレーション失敗: %w", err)
	}
	
	log.Println("DBマイグレーション完了")
	return nil

}