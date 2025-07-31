package registry

import(
    _ "github.com/taatolu/ParkingHub/api/domain/service"
    "github.com/taatolu/ParkingHub/api/domain/repository"
    "github.com/taatolu/ParkingHub/api/infrastructure/infrarepo"
    "github.com/taatolu/ParkingHub/api/usecase"
    _ "github.com/taatolu/ParkingHub/api/presentation/http/handler"
    )
    
type Registry struct{}

func NewRegistry() *Registry {
    return &Registry{}
}

//// --- リポジトリインターフェースと実装（implementation）をつなぐ ---
func (r *Registry) NewCarOwnerRepository() repository.CarOwnerRepository {
    return &infrarepo.CarOwnerRepositoryImpl{}
}


// --- Usecase生成 ---
//Usecaseはリポジトリをラップしているので、Usecaseを作成するためにはリポジトリが必要
func (r *Registry) NewCarOwnerUsecase() usecase.CarOwnerUsecaseIF{
    return usecase.CarOwnerUsecase{carOwnerPero: r.NewCarOwnerRepository()}
}


// --- Handler生成 ---

