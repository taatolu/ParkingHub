package main

import (
	"github.com/taatolu/ParkingHub/api/config"
	presentation "github.com/taatolu/ParkingHub/api/presentation/http"
	"github.com/taatolu/ParkingHub/api/registry"
	"log"
	"net/http"
)

func main() {
	//configの設定
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("設定読み込み失敗:", err)
	}

	// デバッグ用ログ（一時的に追加）
	log.Printf("=== 設定確認 ===")
	log.Printf("DB_HOST: '%s'", conf.DBHost)
	log.Printf("DB_NAME: '%s'", conf.DBName)
	log.Printf("DB_USER: '%s'", conf.DBUser)
	log.Printf("DB_PASSWORD: '%s'", conf.DBPass)
	log.Printf("================")


	reg := registry.NewRegistry()
	defer reg.Close() //アプリ終了時に安全にDBクローズするためにregistry.goに作成したもの

	handler := presentation.InitRouters(reg)

	log.Println("サーバ起動: http://localhost:8080")

	err = http.ListenAndServe(":8080", handler) //作成したマルチプレクサでサーバ起動
	if err != nil {
		log.Fatal(err)
	}
}
