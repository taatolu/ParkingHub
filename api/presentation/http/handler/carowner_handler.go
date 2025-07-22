package handler

import(
    "encoding/json"
    "net/http"
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
    
    //model構築
    owner := &model.CarOwner{
        ID: param.ID,
        FirstName:  param.FirstName,
        MiddleName: param.MiddleName,
        LastName:   param.LastName,
        // 文字列 → time.Time変換する処理が必要
        LicenseExpiration:  param.LicenseExpiration,
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