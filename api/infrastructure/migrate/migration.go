package migrate

import(
	"log"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/taatolu/ParkingHub/api/config"
)

func RunMigration(conf config.Config) error {
	dbHost := conf.DBHost
	dbName := conf.DBName
	dbUser := conf.DBUser
	dbPassword := conf.DBPass
	
	//golang-migrateではPostgreSQL URL 形式を使用する必要があります
	databaseURL := fmt.Sprintf(`postgres://%s:%s@%s:5432/%s?sslmode=disable`, dbUser, dbPassword, dbHost, dbName)

	m, err := migrate.New("file://migrations", databaseURL)
	if err != nil {
		return fmt.Errorf("マイグレーション初期化エラー: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("マイグレーション実行エラー: %w", err)
	}
	
	log.Println("DBマイグレーション完了")
	return nil

}