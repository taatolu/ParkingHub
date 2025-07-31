package registry

import(
    _ "github.com/taatolu/ParkingHub/api/domain/service"
    "github.com/taatolu/ParkingHub/api/domain/repository"
    "github.com/taatolu/ParkingHub/api/infrastructure/infrarepo"
    _ "github.com/taatolu/ParkingHub/api/usecase"
    _ "github.com/taatolu/ParkingHub/api/presentation/http/handler"
    )
    
type Registry struct{}

func NewRegistry() *Registry {
    return &Registry{}
}

//// --- リポジトリインターフェースと実装（implementation）をつなぐ ---
func (r *Registry) NewCarOwnerRepository () repository.CarOwnerRepository {
    return &infrarepo.CarOwnerRepositoryImpl{}
}

// --- Service生成 ---
func (r *Registry) NewCarOwnerService () service.CarOwnerValidation {
    return 
}


// --- Usecase生成 ---


// --- Handler生成 ---

