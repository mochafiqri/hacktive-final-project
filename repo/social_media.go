package repo

import (
	"finalProject/model"
	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	Create(req *model.SocialMedia, trx ...*gorm.DB) error
	Update(req *model.SocialMedia, trx ...*gorm.DB) error
	Delete(req *model.SocialMedia, trx ...*gorm.DB) error
	Get() ([]model.SocialMedia, error)
	GetById(id int) (*model.SocialMedia, error)
}

type SocialMediaContract struct {
	db *gorm.DB
}

func (s *SocialMediaContract) Create(req *model.SocialMedia, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = s.NewDb(trx...)
	var err = newDb.Create(req).Error
	return err
}

func (s *SocialMediaContract) Update(req *model.SocialMedia, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = s.NewDb(trx...)
	var err = newDb.Where("id = ?", req.Id).Updates(req).Error
	return err
}

func (s *SocialMediaContract) Delete(req *model.SocialMedia, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = s.NewDb(trx...)
	var err = newDb.Where("id = ?", req.Id).Delete(req).Error
	return err
}

func (s *SocialMediaContract) Get() ([]model.SocialMedia, error) {
	//TODO implement me
	var result = []model.SocialMedia{}
	var err = s.db.Preload("User").Find(&result).Error
	return result, err
}

func (s *SocialMediaContract) GetById(id int) (*model.SocialMedia, error) {
	//TODO implement me
	var result = model.SocialMedia{}
	var err = s.db.Where("id = ?", id).Find(&result).Error
	return &result, err
}

func (s *SocialMediaContract) NewDb(trx ...*gorm.DB) *gorm.DB {
	if len(trx) > 0 {
		return trx[0]
	}
	return s.db
}
func NewSocialMediaRepo(db *gorm.DB) SocialMediaRepo {
	return &SocialMediaContract{db: db}
}
