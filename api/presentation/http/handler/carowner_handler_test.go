package handler

import(
	"github.com/taatolu/ParkingHub/api/usecase"
	"github.com/taatolu/ParkingHub/api/domain/model"
)

//CarOwneeUsecaseのモック作成
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

// httptest.NewRecorder() でレスポンス記録
// http.NewRequest() でリクエスト作成
// handler.ServeHTTPで実行
// recorder.Result() でレスポンス検証