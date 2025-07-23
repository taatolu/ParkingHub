package http

import (
    "net/http"
    "github.com/taatolu/ParkingHub/api/presentation/http/handler"
    )

func InitRouters()http.Handler{
    http.HandleFunc("/api/v1/car_owners", &handler.CarOwnerHandler{})   //GET(ALL),POST メソッド
    http.HandleFunc("/api/v1/car_owners/", &handler.CarOwnersHandler{})    //GET,PUSH,DELETE メソッド
    return http.DefaultServeMux
}