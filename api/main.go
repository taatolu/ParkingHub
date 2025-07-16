package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("取り急ぎ起動")
	http.ListenAndServe(":8080", nil) //とりあえずデフォルトのマルチプレクサでサーバ起動
}
