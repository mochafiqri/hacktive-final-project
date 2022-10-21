package controller

import (
	"errors"
	"finalProject/helper"
	"finalProject/model"
	"finalProject/parameter"
	"finalProject/repo"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UserController interface {
	Register(req *parameter.RegisterReq) (*parameter.UserResp, int, error)
	Login(req *parameter.LoginReq) (*parameter.LoginResp, int, error)
	Update(req *parameter.UpdateReq) (*parameter.UserResp, int, error)
	Delete(userId int) (int, error)
}

type UserControllerContract struct {
	db *gorm.DB
	repo.UserRepo
}

func (u *UserControllerContract) Delete(userId int) (int, error) {
	//TODO implement me
	user, err := u.UserRepo.FindByField("id", userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user.Id == 0 {
		return http.StatusBadRequest, errors.New(helper.UserNotFound)
	}

	err = u.UserRepo.Delete(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (u *UserControllerContract) Update(req *parameter.UpdateReq) (*parameter.UserResp, int, error) {
	//TODO implement me
	var user, err = u.UserRepo.FindByField("id", req.Id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if user.Id == 0 {
		return nil, http.StatusBadRequest, errors.New(helper.UserNotFound)
	}

	if req.Email != user.Email {
		tmp, err := u.UserRepo.FindByField("email", req.Email)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		if tmp.Id != 0 {
			return nil, http.StatusBadRequest, errors.New(helper.EmailUsed)
		}
		user.Email = req.Email
	}

	if req.Username != user.Username {
		tmp, err := u.UserRepo.FindByField("username", req.Username)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		if tmp.Id != 0 {
			return nil, http.StatusBadRequest, errors.New(helper.UsernameUsed)
		}
		user.Username = req.Username
	}

	err = u.UserRepo.Updates(user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &parameter.UserResp{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}, http.StatusOK, nil
}

func (u *UserControllerContract) Login(req *parameter.LoginReq) (*parameter.LoginResp, int, error) {
	//TODO implement me
	user, err := u.UserRepo.FindByField("email", req.Email)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if user.Id == 0 {
		return nil, http.StatusBadRequest, errors.New(helper.UserNotFound)
	}

	err = helper.ValidatePassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, http.StatusBadRequest, errors.New(helper.PasswordInvalid)
	}

	token, err := helper.GenerateToken(&helper.Token{
		UserId:  user.Id,
		Email:   user.Email,
		Expired: time.Time{},
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &parameter.LoginResp{Token: token}, http.StatusOK, nil

}

func (u *UserControllerContract) Register(req *parameter.RegisterReq) (*parameter.UserResp, int, error) {
	//check email
	tmpUser, err := u.UserRepo.FindByField("email", req.Email)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if tmpUser.Id != 0 {
		return nil, http.StatusBadRequest, errors.New(helper.EmailUsed)
	}

	tmpUser, err = u.UserRepo.FindByField("username", req.Username)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if tmpUser.Id != 0 {
		return nil, http.StatusBadRequest, errors.New(helper.UsernameUsed)
	}

	hash, err := helper.GeneratePassword([]byte(req.Password))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	var user = model.User{
		Id:        0,
		Username:  req.Username,
		Email:     req.Email,
		Password:  hash,
		Age:       req.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.UserRepo.Register(&user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &parameter.UserResp{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}, http.StatusCreated, nil
}

func NewUserController(db *gorm.DB) UserController {
	return &UserControllerContract{
		db:       db,
		UserRepo: repo.NewUserRepo(db),
	}
}
