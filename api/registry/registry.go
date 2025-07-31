package registry

import(
    "github.com/taatolu/ParkingHub/domain/service"
    "github.com/taatolu/ParkingHub/domain/repository"
    "github.com/taatolu/ParkingHub/infrastructure/infrarepo"
    "github.com/taatolu/ParkingHub/usecase"
    "github.com/taatolu/ParkingHub/presentation/http/handler"
    )
    
type Registry struct{}

func NewRegistry() *Registry {
    return &Registry{}
}

//// --- リポジトリインターフェースと実装（implementation）をつなぐ ---
func (r *Registry) NewCarOwnerRepository () repository.CarOwnerRepository {
    return infrarepo.CarOwnerRepositoryImpl
}
