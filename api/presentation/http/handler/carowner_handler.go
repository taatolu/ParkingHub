package handler

import (
	"encoding/json"
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/usecase"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CarOwnerHandler struct {
	//usecase層のインターフェースから実装
	Usecase usecase.CarOwnerUsecaseIF
}

// CarOwnerHandler definition（ルーターでCarOwnerHandlerが呼ばれたときどのメソッドを実行するか & ServeHTTPをラップ）
func (h CarOwnerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/api/v1/car_owners" && r.Method == http.MethodPost:
		h.CreateCarOwner(w, r)
	case strings.HasPrefix(r.URL.Path, "/api/v1/car_owners/") && r.Method == http.MethodGet:
		h.FindByID(w, r)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf(`{"error":"リクエストメソッドが不正です"}`)))
	}
}

// POST api car_owners
func (h CarOwnerHandler) CreateCarOwner(w http.ResponseWriter, r *http.Request) {
	//リクエストボディの内容を取得
	///取得したリクエストボディの内容を格納する構造体を作成
	var param struct {
		ID                string `json:"id"`
		FirstName         string `json:"first_name"`
		MiddleName        string `json:"middle_name"`
		LastName          string `json:"last_name"`
		LicenseExpiration string `json:"license_expiration"`
	}
	///リクエストボディの内容をparamにパース
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//取得したリクエストボディの型（取得時は文字列）をエンティティの型と一致するよう修正
	idInt, err := strconv.Atoi(param.ID)
	if err != nil {
		http.Error(w, "IDの型変換に失敗", http.StatusBadRequest)
		return
	}

	expiry, err := time.Parse("2006-01-02", param.LicenseExpiration)
	if err != nil {
		http.Error(w, "Invalid LicenseExpiration format", http.StatusBadRequest)
		return
	}

	//model構築
	owner := &model.CarOwner{
		ID:         idInt,
		FirstName:  param.FirstName,
		MiddleName: param.MiddleName,
		LastName:   param.LastName,
		// 文字列 → time.Time変換する処理が必要
		LicenseExpiration: expiry,
	}

	// ユースケースを呼んで新規登録
	err = h.Usecase.RegistCarOwner(owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(owner)
}

// TODO: handlerを本実装に差し替える（Issue #54）
func (h CarOwnerHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
    
    //パス形式の検証: /api/v1/car_owners/{id}
    if !strings.HasPrefix(path, "/api/v1/car_owners/"){
        http.Error(w, "パスが不正です", http.StatusBadRequest)
        return
    }
	//パラメータを取得（パスから"/api/v1/car_owners/"を引く）
	idStr := strings.TrimPrefix(path, "/api/v1/car_owners/")
	if idStr == "" {
	    http.Error(w, "パラメータが存在しません", http.StatusBadRequest)
	    return
	}

	//パスパラメータが数値に変換できない場合
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error:パスパラメータが数値でありません", http.StatusBadRequest)
		return
	}

	owner, err := h.Usecase.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// CarOwnerの構造体を返す
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err = json.NewEncoder(w).Encode(owner); err != nil {
        http.Error(w, "error:エンコード失敗", http.StatusInternalServerError)
    }

}
