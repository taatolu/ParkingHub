package handler

import(
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/usecase"
	"testing"
	"net/http/httptest"
	"net/http"
)

//CarOwneeUsecaseのモック作成
///暗黙的にインターフェース宣言…Goでは明示的にinterfaseを宣言しなくても、同じメソッドを持っていれば「満たしている」ことになる
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
    
    //ハンドラーのインスタンス生成時にmockUsecaseを型アサーションする
    ///上記Usecaseのモック（MockCarOwnerUsecase）を暗黙的に宣言した為の弊害（型アサーションしないと認識してくれない）
    handler := &CarOwnerHandler{Usecase: usecase.CarOwnerUsecase(mockUsecase)}
    
    // httptest.NewRecorder() でレスポンス記録
    rec := httptest.NewRecorder()
    
    // http.NewRequest() でリクエスト作成
    req, err := http.NewRequest("POST", "/api/v1/car_owners", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    // handler.ServeHTTPで実行
    handler.ServeHTTP(rec, req)
    
    // recorder.Result() でレスポンス検証
    resp := rec.Result()
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusCreated {
        t.Errorf("GotStatus=%d WantStatus=%d", resp.StatusCode, http.StatusCreated)
    }

}

