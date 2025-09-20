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
		//パラメータの値を取得
		param := strings.TrimPrefix(r.URL.Path, "/api/v1/car_owners/")
		//パラメータの値が数値か文字列かでパンドラの呼び分け
		if _, err := strconv.Atoi(param); err != nil {
			//取得したパラメーターが数値でない場合
			h.FindByName(w, r)
		} else {
			//取得したパラメーターが数値の場合
			h.FindByID(w, r)
		}
	case strings.HasPrefix(r.URL.Path, "/api/v1/car_owners/") && r.Method == http.MethodPut:
		h.Update(w, r)

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
	idUint := uint(idInt)

	expiry, err := time.Parse("2006-01-02", param.LicenseExpiration)
	if err != nil {
		http.Error(w, "Invalid LicenseExpiration format", http.StatusBadRequest)
		return
	}

	//model構築
	owner := &model.CarOwner{
		ID:         idUint,
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

// GET (Find By id)
func (h CarOwnerHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
    
    //パス形式の検証: /api/v1/car_owners/{id}
    if !strings.HasPrefix(path, "/api/v1/car_owners/"){
        http.Error(w, "error:パスが不正です", http.StatusBadRequest)
        return
    }
	//パラメータを取得（パスから"/api/v1/car_owners/"を引く）
	idStr := strings.TrimPrefix(path, "/api/v1/car_owners/")
	if idStr == "" {
	    http.Error(w, "error:パラメータが存在しません", http.StatusBadRequest)
	    return
	}

	//パスパラメータが数値に変換できない場合
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error:パスパラメータが数値でありません", http.StatusBadRequest)
		return
	}

	uid := uint(id)

	//Usecaeに渡して検索してもらう
	owner, err := h.Usecase.FindByID(uid)
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

// FindByNameは、URLパスから名前で車の所有者を取得するGETリクエストを処理します。
func (h *CarOwnerHandler) FindByName (w http.ResponseWriter, r *http.Request) {
	//pathの検証
	path := r.URL.Path
	if !strings.HasPrefix(path, "/api/v1/car_owners/") {
		http.Error(w, "error:Pathが不正です", http.StatusBadRequest)
		return
	}

	//パラメーターを取得
	name := strings.TrimPrefix(path, "/api/v1/car_owners/")
	if name == "" {
		http.Error(w, "error:パラメータが存在しません", http.StatusBadRequest)
		return
	}

	//Usecaseに渡して検索してもらう
	owners, err := h.Usecase.FindByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	//CarOwner構造体のリストを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(owners)
	if err != nil {
		http.Error(w, "error:エンコード失敗", http.StatusInternalServerError)
	}
}

//Updateメソッドは、車の所有者の情報を更新するためのHTTPメソッドPUTリクエストを処理します。
func (h *CarOwnerHandler) Update (w http.ResponseWriter, r *http.Request){
	//メソッドの判定
	if r.Method != http.MethodPut {
        // クライアントが不正なHTTPメソッドでアクセスした場合
        http.Error(w, `{"error":"リクエストメソッドが不正です"}`, http.StatusMethodNotAllowed)
        return
    }

	//URLPathの検証
	path := r.URL.Path
	//HasPrefixの第2引数の値がpathの先頭にあるかどうかをチェック
	if !strings.HasPrefix(path, "/api/v1/car_owners/") {
		http.Error(w, "error:Pathが不正です", http.StatusBadRequest)
		return
	}

	//Pathからパラメーターを取得
	idStr := strings.TrimPrefix(path, "/api/v1/car_owners/")
	//パラメーターが存在しない場合
	if idStr == "" {
		http.Error(w, "error: パラメータが存在しません", http.StatusBadRequest)
		return
	}

	//パスパラメーターを数値に変換
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error:パスパラメーターが数値ではありません", http.StatusBadRequest)
		return
	}
	uid := uint(id)		//uint型に変換

	//RequestBodyの内容を取得(UsecaseのUpdateメソッドに渡す引数の準備)
	//取得したBodyの内容を保存するための構造体を作成(Bodyの内容はstringで取得される)
	var param struct {
		//IDについてはパスパラメータの値を使用するのでBodyの値を取得しない
		FirstName	string `json:"first_name"`
		MiddleName	string `json:"middle_name"`
		LastName	string `json:"last_name"`
		LicenseExpiration	string `json:"license_expiration"`
	}

	//requestBodyの内容をparamにパースする
	err = json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		http.Error(w, "error:Bodyの内容が不正です", http.StatusBadRequest)
		return
	}

	//Updateメソッドに渡せるようにするために*modei.CarOwnerの値を作成
	//その前に、日付を扱うフィールド値を調整
	expiry, err := time.Parse("2006-01-02", param.LicenseExpiration)
	if err != nil {
		http.Error(w, "error:Bodyから取得したLicenseExpirationを日付型に変更できない", http.StatusBadRequest)
		return
	}

	//modelの作成
	owner := &model.CarOwner{
		ID:					uid,
		FirstName:			param.FirstName,
		MiddleName:			param.MiddleName,
		LastName:			param.LastName,
		LicenseExpiration:	expiry,
	}

	//UsecaseのUpdateメソッドを読んで更新
	err = h.Usecase.Update(owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Updateが成功したら
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(owner)
}



