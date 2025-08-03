package registry

import(
    "database/sql"
    "github.com/taatolu/ParkingHub/api/config"
    "github.com/taatolu/ParkingHub/api/infrastructure/postgres"
    _ "github.com/taatolu/ParkingHub/api/domain/service"
    "github.com/taatolu/ParkingHub/api/domain/repository"
    "github.com/taatolu/ParkingHub/api/infrastructure/infrarepo"
    "github.com/taatolu/ParkingHub/api/usecase"
    "github.com/taatolu/ParkingHub/api/presentation/http/handler"
    )
    
type Registry struct{
    //DBフィールドを作成
    DB *sql.DB
}

//Registryのファクトリ関数
func NewRegistry() *Registry {
    //Registryのファクトリ関数の中で、DBにpostgresをDIする
    ///configで環境変数取得
    conf, err := config.LoadConfig()
    if err != nil {
        panic(fmt.Errorf("環境変数の取得失敗; %w",err))
    }
    ///上記環境変数をもとに、postgresのイニシャライズ
    db, err := postgres.InitPostgres(conf)
    if err != nil {
        panic(fmt.Errorf("DBの初期化失敗; %w", err))
    }
    //DBを依存注入したRegistryを返却
    return &Registry{DB: db}
}

//// --- リポジトリインターフェースと実装（implementation）をつなぐ ---
func (r *Registry) NewCarOwnerRepository() repository.CarOwnerRepository {
    //CarOwnerRepositoryImplはDBフィールドを持つので、RegistryにDIしたpostgrtesをさらにDI
    return &infrarepo.CarOwnerRepositoryImpl{DB: r.DB}
}


// --- Usecase生成 ---
//Usecaseはリポジトリをラップしているので、Usecaseを作成するためにはリポジトリが必要
func (r *Registry) NewCarOwnerUsecase() usecase.CarOwnerUsecaseIF {
    return &usecase.CarOwnerUsecase{CarOwnerRepo: r.NewCarOwnerRepository()}
}


// --- Handler生成 ---
func (r *Registry) NewCarOwnerHandler() handler.CarOwnerHandler {

}
