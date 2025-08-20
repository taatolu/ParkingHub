package http

import (
    "net/http"
    "github.com/taatolu/ParkingHub/api/registry"
    _ "github.com/taatolu/ParkingHub/api/presentation/http/handler"
    )

func InitRouters(reg *registry.Registry)http.Handler{
    //マルチプレクサを作成
    mux := http.NewServeMux()
    mux.Handle("/api/v1/car_owners", reg.NewCarOwnerHandler())   //GET(ALL),POST メソッド
    http.Handle("/api/v1/car_owners/", &handler.CarOwnersHandler{})    //GET,PUSH,DELETE メソッド
    return mux
}