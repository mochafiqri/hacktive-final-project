package repo

import (
	"finalProject/model"
	"gorm.io/gorm"
)

type CommentRepo interface {
	Create(req *model.Comment, trx ...*gorm.DB) error
	Update(req *model.Comment, trx ...*gorm.DB) error
	Delete(req *model.Comment, trx ...*gorm.DB) error
	Get() ([]model.Comment, error)
	GetById(id int) (*model.Comment, error)
}

type CommentRepoContract struct {
	db *gorm.DB
}

func (c *CommentRepoContract) Update(req *model.Comment, trx ...*gorm.DB) error {
	var newDb = c.NewDb(trx...)

	var err = newDb.Where("id = ? ", req.Id).Updates(req).Error
	return err
}

func (c *CommentRepoContract) Delete(req *model.Comment, trx ...*gorm.DB) error {
	var newDb = c.NewDb(trx...)

	var err = newDb.Where("id = ? ", req.Id).Delete(req).Error
	return err
}

func (c *CommentRepoContract) GetById(id int) (*model.Comment, error) {
	var result = model.Comment{}
	var err = c.db.Where("id = ? ", id).Find(&result).Error
	return &result, err
}

func (c *CommentRepoContract) Get() ([]model.Comment, error) {
	//TODO implement me
	var result = []model.Comment{}

	var err = c.db.
		Preload("User").
		Preload("Photo").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CommentRepoContract) Create(req *model.Comment, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = c.NewDb(trx...)
	var err = newDb.Create(req).Error
	return err
}

func (c *CommentRepoContract) NewDb(trx ...*gorm.DB) *gorm.DB {
	if len(trx) > 0 {
		return trx[0]
	}
	return c.db
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &CommentRepoContract{db: db}
}
