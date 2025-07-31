package registry

import(
    "github.com/taatolu/ParkingHub/api/domain/service"
    "github.com/taatolu/ParkingHub/api/domain/repository"
    "github.com/taatolu/ParkingHub/api/infrastructure/infrarepo"
    "github.com/taatolu/ParkingHub/api/usecase"
    "github.com/taatolu/ParkingHub/api/presentation/http/handler"
    )
    
type Registry struct{}

func NewRegistry() *Registry {
    return &Registry{}
}

//// --- リポジトリインターフェースと実装（implementation）をつなぐ ---
func (r *Registry) NewCarOwnerRepository () repository.CarOwnerRepository {
    return &infrarepo.CarOwnerRepositoryImpl{}
}
