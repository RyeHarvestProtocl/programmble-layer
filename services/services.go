package services

import (
	"gorm.io/gorm"
)

type BaseService[T any] struct {
	Db *gorm.DB
}

func NewBaseService[T any](Db *gorm.DB) *BaseService[T] {
	return &BaseService[T]{Db: Db}
}

func (service *BaseService[T]) Create(item *T) error {
	return service.Db.Create(item).Error
}

func (service *BaseService[T]) GetAll() ([]T, error) {
	var items []T
	err := service.Db.Find(&items).Error
	return items, err
}

func (service *BaseService[T]) GetById(id uint) (*T, error) {
	var item T
	err := service.Db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (service *BaseService[T]) Find(conds ...interface{}) ([]T, error) {
	var items []T
	if len(conds) >= 1 {
		err := service.Db.Where(conds[0], conds[1:]...).Find(&items).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := service.Db.Find(&items).Error
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}

func (service *BaseService[T]) First(dest *T, conds ...interface{}) error {
	if len(conds) >= 1 {
		err := service.Db.Where(conds[0], conds[1:]...).First(dest).Error
		if err != nil {
			return err
		}
	} else {
		err := service.Db.First(dest).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *BaseService[T]) Last(dest *T, conds ...interface{}) error {
	if len(conds) >= 1 {
		err := service.Db.Where(conds[0], conds[1:]...).Last(dest).Error
		if err != nil {
			return err
		}
	} else {
		err := service.Db.Last(dest).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *BaseService[T]) Update(getModelInstance func() T, conds interface{}, updates interface{}) error {
	modelInstance := getModelInstance()
	result := service.Db.Model(modelInstance).Where(conds).Updates(updates)
	return result.Error
}
