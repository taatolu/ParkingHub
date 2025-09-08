package registry

import(
    "database/sql"
    "fmt"
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
    DB *gorm.DB
}

// --- DBクローズ用の設定 ---
func (r *Registry) Close() error {
    //アプリケーション終了時にDBを安全にクローズするためのメソッド
    if r.DB != nil {
        return r.DB.Close()
    }
    return nil
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
    //*返り値がインターフェースの場合、そのインターフェースを満たす値型でもポインタ型でもどちらでも返せます。（repository.CarOwnerRepositoryはインターフェース型）
}


// --- Usecase生成 ---
//Usecaseはリポジトリをラップしているので、Usecaseを作成するためにはリポジトリが必要
func (r *Registry) NewCarOwnerUsecase() usecase.CarOwnerUsecaseIF {
    return &usecase.CarOwnerUsecase{CarOwnerRepo: r.NewCarOwnerRepository()}
    //*返り値がインターフェースの場合、そのインターフェースを満たす値型でもポインタ型でもどちらでも返せます。（usecase.CarOwnerUsecaseIFはインターフェース型）
}


// --- Handler生成 ---
func (r *Registry) NewCarOwnerHandler() *handler.CarOwnerHandler {
    //*返り値が構造体の場合、ポインタ型は返せません（handler.CarOwnersHandlerは構造体なので、”*”をつけてポインタで返すようにした）
    return &handler.CarOwnerHandler{
        Usecase: r.NewCarOwnerUsecase(),
    }
}


// --- ロギングのDI ---
// --- ルーターの初期化をここで行ってもOK --- 