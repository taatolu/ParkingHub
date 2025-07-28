package handler

import(
	"github.com/taatolu/ParkingHub/api/domain/model"
	"testing"
	"net/http/httptest"
	"net/http"
)

//CarOwneeUsecaseのモック作成
///同じメソッドを持っていればインターフェースを「満たしている」ことになる
type MockCarOwnerUsecase struct{
	//usecase/carowner_handler.goのcarOwnerのメソッドセットを満たす
	RegistCarOwnerFunc func(owner *model.CarOwner) error
}

func (m *MockCarOwnerUsecase) RegistCarOwner(owner *model.CarOwner) error {
	if m.RegistCarOwnerFunc != nil{
		//RegistCarOwnerFuncがnilでなかったら
		return m.RegistCarOwnerFunc(owner)
	}
	return nil
}

//テストの実行
func TestRegistCarOwner(t *testing.T){
    //CarOwnerUsecaseモックのインスタンス化
    mockUsecase := &MockCarOwnerUsecase{
        RegistCarOwnerFunc : func(owner *model.CarOwner) error{
            return nil
        },
    }
    
    //ハンドラーのインスタンス生成
    handler := &CarOwnerHandler{Usecase: mockUsecase}
    
    // httptest.NewRecorder() でレスポンス記録
    rec := httptest.NewRecorder()
    
    // http.NewRequest() でリクエスト作成
    ///http.NewRequestの第3引数に渡すrequest.Body(json)を作成
    body := `{"id":"1", "first_name":"test", "middle_name":"山田", "last_name":"太郎", "license_expiration":"1234"}`
    req, err := http.NewRequest("POST", "/api/v1/car_owners", strings.NewReader(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    
    // handler.ServeHTTPで実行
    handler.ServeHTTP(rec, req)
    
    // recorder.Result() でレスポンス検証
    resp := rec.Result()
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusCreated {
        t.Errorf("GotStatus=%d WantStatus=%d", resp.StatusCode, http.StatusCreated)
    }
    
}

