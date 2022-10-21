package repo

import (
	"finalProject/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindByField(by string, value interface{}) (*model.User, error)
	Register(user *model.User, trx ...*gorm.DB) error
	Updates(user *model.User, trx ...*gorm.DB) error
	Delete(user *model.User, trx ...*gorm.DB) error
}

type UserRepoContract struct {
	db *gorm.DB
}

func (u *UserRepoContract) Delete(user *model.User, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = u.newDb(trx...)

	var err = newDb.Where("id = ?", user.Id).Delete(user).Error
	return err
}

func (u *UserRepoContract) Updates(user *model.User, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = u.newDb(trx...)

	var err = newDb.Updates(user).Error
	return err
}

func (u *UserRepoContract) Register(user *model.User, trx ...*gorm.DB) error {
	//TODO implement me
	var newDb = u.newDb(trx...)

	var err = newDb.Create(user).Error
	return err
}

func (u *UserRepoContract) FindByField(by string, value interface{}) (*model.User, error) {
	//TODO implement me
	var result = model.User{}
	var query = by + "= ?"
	var err = u.db.Where(query, value).Debug().
		Find(&result).Error

	return &result, err
}

func (u *UserRepoContract) newDb(trx ...*gorm.DB) *gorm.DB {
	var newDb *gorm.DB
	if trx != nil {
		newDb = trx[0]
	} else {
		newDb = u.db
	}
	return newDb
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoContract{db: db}
}
