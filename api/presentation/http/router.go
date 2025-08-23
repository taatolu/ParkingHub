package http

import (
	"github.com/taatolu/ParkingHub/api/presentation/http/handler"
	"github.com/taatolu/ParkingHub/api/registry"
	"net/http"
)

func InitRouters(reg *registry.Registry) http.Handler {
	//マルチプレクサを作成
	mux := http.NewServeMux()

	carOwnerHandler := reg.NewCarOwnerHandler()

	mux.Handle("/api/v1/car_owners", carOwnerHandler)   //GET(ALL),POST メソッド
	mux.Handle("/api/v1/car_owners/", carOwnerHandler) //GET,PUSH,DELETE メソッド
	return mux
}
