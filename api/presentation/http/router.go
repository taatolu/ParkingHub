package http

import (
    "net/http"
    )

func InitRouters()http.Handler{
    http.HandleFunc("/api/v1/car_owners", *handler.CarOwnerHandler)   //GET(ALL),POST メソッド
    http.HandleFunc("/api/v1/car_owners/", *handler.CarOwnersHandler)    //GET,PUSH,DELETE メソッド
    return http.DefaultServeMux
}