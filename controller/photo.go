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

type PhotoController interface {
	Add(req *parameter.PhotoReq) (*parameter.PhotoResp, int, error)
	Get() ([]parameter.PhotoListResp, int, error)
	Update(req *parameter.PhotoReq) (*parameter.PhotoResp, int, error)
	Delete(req *parameter.PhotoReq) (int, error)
}

type PhotoControllerContract struct {
	db        *gorm.DB
	photoRepo repo.PhotoRepo
	userRepo  repo.UserRepo
}

func (p *PhotoControllerContract) Update(req *parameter.PhotoReq) (*parameter.PhotoResp, int, error) {
	//TODO implement me
	var photo, err = p.photoRepo.Get(req.Id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if photo.Id == 0 {
		return nil, http.StatusNotFound, errors.New(helper.PhotoNotFound)
	}

	if photo.UserId != req.UserId {
		return nil, http.StatusUnauthorized, err
	}

	photo.Title = req.Title
	photo.Caption = req.Caption
	photo.PhotoUrl = req.PhotoUrl

	err = p.photoRepo.Updates(photo)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &parameter.PhotoResp{
		Id:       photo.Id,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserId:   photo.UserId,
		Created:  photo.CreatedAt,
	}, http.StatusOK, nil

}

func (p *PhotoControllerContract) Delete(req *parameter.PhotoReq) (int, error) {
	//TODO implement me
	var photo, err = p.photoRepo.Get(req.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if photo.Id == 0 {
		return http.StatusNotFound, errors.New(helper.PhotoNotFound)
	}

	if photo.UserId != req.UserId {
		return http.StatusUnauthorized, err
	}

	err = p.photoRepo.Delete(&model.Photo{
		Id: photo.Id,
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil

}

func (p *PhotoControllerContract) Get() ([]parameter.PhotoListResp, int, error) {
	//TODO implement me
	photo, err := p.photoRepo.List()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(photo) == 0 {
		return nil, http.StatusNotFound, errors.New(helper.PhotoNotFound)
	}
	var resp = []parameter.PhotoListResp{}
	for _, v := range photo {
		var tmp = parameter.PhotoListResp{
			Id:        v.Id,
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoUrl:  v.PhotoUrl,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		tmp.User.Email = v.User.Email
		tmp.User.Username = v.User.Username
		resp = append(resp, tmp)
	}

	return resp, http.StatusOK, nil
}

func (p *PhotoControllerContract) Add(req *parameter.PhotoReq) (*parameter.PhotoResp, int, error) {

	user, err := p.userRepo.FindByField("id", req.UserId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if user.Id == 0 {
		return nil, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized))
	}

	var photo = model.Photo{
		Title:     req.Title,
		Caption:   req.Caption,
		PhotoUrl:  req.PhotoUrl,
		UserId:    req.UserId,
		CreatedAt: time.Now(),
	}
	err = p.photoRepo.Create(&photo)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &parameter.PhotoResp{
		Id:       photo.Id,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserId:   photo.UserId,
		Created:  photo.CreatedAt,
	}, http.StatusCreated, nil
}

func NewPhotoController(db *gorm.DB) PhotoController {
	return &PhotoControllerContract{
		db:        db,
		photoRepo: repo.NewPhotoRepo(db),
		userRepo:  repo.NewUserRepo(db),
	}
}
