package main

import (
	"log"
	"net/http"
	"github.com/taatolu/ParkingHub/api/config"
	"github.com/taatolu/ParkingHub/api/registry"
	"github.com/taatolu/ParkingHub/api/infrastructure/migrate"
	presentation "github.com/taatolu/ParkingHub/api/presentation/http"
)

func main() {
	//configの設定
	conf, err := config.LoadConfig()
	if err != nil {
        log.Fatal("設定読み込み失敗:", err)
    }

	//DBbのマイグレーション
	if err := migrate.RunMigration(*conf); err != nil {
	log.Fatal("マイグレーション失敗:", err)
    }

	reg := registry.NewRegistry()
	defer reg.Close()	//アプリ終了時に安全にDBクローズするためにregistry.goに作成したもの

	router := presentation.InitRouters(reg)

	log.Println("サーバ起動: http://localhost:8080")
	err := http.ListenAndServe(":8080", router) //作成したマルチプレクサでサーバ起動
	if err != nil{
		log.Fatal(err)
	}
}
