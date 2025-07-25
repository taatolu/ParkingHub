package handler

import(
    "strconv"
    "encoding/json"
    "net/http"
    "time"
    "github.com/taatolu/ParkingHub/api/usecase"
    "github.com/taatolu/ParkingHub/api/domain/model"
    )

type CarOwnerHandler struct{
    Usecase usecase.CarOwnerUsecase
}

//POST api car_owners
func (h CarOwnerHandler) CreateCarOwner(w http.ResponseWriter, r *http.Request){
    //リクエストボディの内容を取得
    ///取得したリクエストボディの内容を格納する構造体を作成
    var param struct{
        ID          string  `json:"id"`
        FirstName   string  `json:"first_name"`
        MiddleName  string  `json:"middle_name"`
        LastName    string  `json:"last_name"`
        LicenseExpiration   string  `json:"license_expiration"`
    }
    ///リクエストボディの内容をparamにパース
    err := json.NewDecoder(r.Body).Decode(&param)
    if err != nil{
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    //取得したリクエストボディの型（取得時は文字列）をエンティティの型と一致するよう修正
    idInt, err := strconv.Atoi(param.ID)
    if err != nil{
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
        ID: idInt,
        FirstName:  param.FirstName,
        MiddleName: param.MiddleName,
        LastName:   param.LastName,
        // 文字列 → time.Time変換する処理が必要
        LicenseExpiration:  expiry,
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