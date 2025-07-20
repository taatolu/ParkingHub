package postgres

import(
	"database/sql"
	"github.com/taatolu/ParkingHub/api/config"
	_ "github.com/lib/pq"
	"fmt"
)

//PostgresSQLへの接続初期化
func InitPostgres(conf *config.Config)(*sql.DB, error){
	connSrt := fmt.Sprintf(`host=%s dbname=%s user=%s password=%s sslmode=disable`,
		conf.DBHost, conf.DBName, conf.DBUser, conf.DBPass)
	//postgreSQLへ接続
	db, err := sql.Open("postgres", connSrt)
	if err != nil{
		return nil, fmt.Errorf("postgresへの接続失敗: %w", err)
	}
	return db, nil
}
