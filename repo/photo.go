package repo

import (
	"finalProject/model"
	"gorm.io/gorm"
)

type PhotoRepo interface {
	List() ([]model.Photo, error)
	GetByUser(userId int) ([]model.Photo, error)
	Get(id int) (*model.Photo, error)
	Create(photo *model.Photo, trx ...*gorm.DB) error
	Updates(photo *model.Photo, trx ...*gorm.DB) error
	Delete(photo *model.Photo, trx ...*gorm.DB) error
}

type PhotoRepoContract struct {
	db *gorm.DB
}

func (p *PhotoRepoContract) List() ([]model.Photo, error) {
	//TODO implement me
	var result = []model.Photo{}
	var err = p.db.Preload("User").Find(&result).Error
	return result, err
}

func (p *PhotoRepoContract) GetByUser(userId int) ([]model.Photo, error) {
	//TODO implement me
	var result = []model.Photo{}
	var err = p.db.Where("user_id = ?", userId).Find(&result).Error
	return result, err
}

func (p *PhotoRepoContract) Get(id int) (*model.Photo, error) {
	//TODO implement me
	var result = model.Photo{}
	var err = p.db.Where("id = ?", id).Find(&result).Error
	return &result, err
}

func (p *PhotoRepoContract) Create(photo *model.Photo, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = p.newDb(trx...)

	var err = newDb.Create(photo).Error
	return err
}

func (p *PhotoRepoContract) Updates(photo *model.Photo, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = p.newDb(trx...)

	var err = newDb.Where("id = ?", photo.Id).
		Updates(photo).Error
	return err
}

func (p *PhotoRepoContract) Delete(photo *model.Photo, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = p.newDb(trx...)

	var err = newDb.Where("id = ?", photo.Id).
		Delete(photo).Error
	return err
}

func (p *PhotoRepoContract) newDb(trx ...*gorm.DB) *gorm.DB {
	if len(trx) > 0 {
		return trx[0]
	}
	return p.db
}

func NewPhotoRepo(db *gorm.DB) PhotoRepo {
	return &PhotoRepoContract{db: db}
}
