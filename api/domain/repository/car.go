package repository

import(
	"github.com/taatolu/ParkingHub/api/domain/model"
)

type CarRepository interface {
	Create(car *model.Car) (*model.Car, error)	// 新しい車をデータベースに保存するためのメソッドです。引数としてCar構造体のポインタを受け取り、保存されたCar構造体のポインタとエラーを返す。
	GetAll() ([]*model.Car, error)	// データベースからすべての車を取得するためのメソッドです。Car構造体のポインタのスライスとエラーを返す。
	GetAllByOwnerID( carOwnerID uint) ([]*model.Car, error)	// 指定された車の所有者IDに関連するすべての車をデータベースから取得するためのメソッドです。Car構造体のポインタのスライスとエラーを返す。
	Get(carID uint) (*model.Car, error)	// 指定された車IDに関連する車をデータベースから取得するためのメソッドです。Car構造体のポインタとエラーを返す。
	Update(carID uint) (*model.Car, error) // 既存の車の情報を更新するためのメソッドです。引数としてCar構造体のポインタを受け取り、更新されたCar構造体のポインタとエラーを返す。
	Delete(carID uint) error	// 指定された車IDに関連する車をデータベースから削除するためのメソッドです。エラーを返す。
}