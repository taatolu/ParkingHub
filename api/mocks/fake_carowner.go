package mocks

import(
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
)

type FakeCarOwnerRepo struct{
	SavedOwner	*model.CarOwner
	AllOwners	[]*model.CarOwner
	WantError	bool	
}

func (f *FakeCarOwnerRepo) Save (carOwner *model.CarOwner) error {
	if carOwner != nil {
		f.SavedOwner = carOwner		//fakerepoの実装で期待通りのcarOwnerが保存されたか確認
		return nil
	}
	return fmt.Errorf("fakerepoImplに渡したデータが不正です")
}

func (f *FakeCarOwnerRepo) GetAll() ([]*model.CarOwner, error) {
	//取り急ぎ、実装しない（未実装であることがわかるようにエラーを返す）
	return nil, fmt.Errorf("fakerepoのGetAllメソッドは未実装です")
}


func (f *FakeCarOwnerRepo) FindByID(id uint) (*model.CarOwner, error) {
	tempOwner := &model.CarOwner{ID:id}
	if !tempOwner.IsIDPositive(){
		return nil, fmt.Errorf("IDが不正です(負の数): %v", id)
	}
    return tempOwner, nil
}

func (f *FakeCarOwnerRepo) FindByName(name string)([]*model.CarOwner, error) {
	foundOwners := []*model.CarOwner{}

	for _, owner := range f.AllOwners {
		if owner.IsContainsName(name) {
			foundOwners = append(foundOwners, owner)
		}
	}
	return foundOwners, nil
}

//Update
//Usecaseテストで呼ばれたときにRepositoryImplがあたかもDBをUpdateしたかのようにふるまう
func (f *FakeCarOwnerRepo) Update (carOwner *model.CarOwner) error {
    if f.SavedOwner == nil {
        return fmt.Errorf("データが存在しません")
    }

	if f.WantError {
		return fmt.Errorf("データの取得でErrorが発生")
	}
    return nil
}


func (f *FakeCarOwnerRepo) Delete(id uint) error {
	//取り急ぎ、実装しない（未実装であることがわかるようにエラーを返す）
	return fmt.Errorf("Deleteメソッドは未実装です")
}