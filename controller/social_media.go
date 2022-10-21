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

type SocialMediaController interface {
	Add(req *parameter.SocialMediaReq) (*parameter.SocialMediaResp, int, error)
	Get() ([]parameter.SocialMediaList, int, error)
	Update(req *parameter.SocialMediaReq) (*parameter.SocialMediaResp, int, error)
	Delete(req *parameter.SocialMediaReq) (int, error)
}

type SocialMediaControllerContract struct {
	db              *gorm.DB
	socialMediaRepo repo.SocialMediaRepo
}

func (s SocialMediaControllerContract) Add(req *parameter.SocialMediaReq) (*parameter.SocialMediaResp, int, error) {
	//TODO implement me
	var socialMedia = model.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserId:         req.UserId,
		CreatedAt:      time.Now(),
	}
	var err = s.socialMediaRepo.Create(&socialMedia)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var resp = parameter.SocialMediaResp{
		Id:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		CreatedAt:      socialMedia.CreatedAt,
	}
	return &resp, http.StatusOK, nil
}

func (s SocialMediaControllerContract) Get() ([]parameter.SocialMediaList, int, error) {
	//TODO implement me
	var resp = []parameter.SocialMediaList{}
	socialMedias, err := s.socialMediaRepo.Get()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(socialMedias) == 0 {
		return nil, http.StatusNotFound, errors.New(helper.CommentNotFound)
	}

	for _, v := range socialMedias {
		var tmp = parameter.SocialMediaList{
			Id:             v.Id,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserId:         v.UserId,
			CreatedAt:      v.CreatedAt,
			User: parameter.UserResp{
				Id:       v.User.Id,
				Username: v.User.Username,
				Email:    v.User.Email,
			},
		}

		resp = append(resp, tmp)
	}
	return resp, http.StatusOK, nil
}

func (s SocialMediaControllerContract) Update(req *parameter.SocialMediaReq) (*parameter.SocialMediaResp, int, error) {
	//TODO implement me
	sm, err := s.socialMediaRepo.GetById(req.Id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if sm.Id == 0 {
		return nil, http.StatusBadRequest, errors.New(helper.CommentNotFound)
	}

	if sm.UserId != req.UserId {
		return nil, http.StatusUnauthorized, errors.New(helper.CommentNotFound)

	}
	sm.Name = req.Name
	sm.SocialMediaUrl = req.SocialMediaUrl

	err = s.socialMediaRepo.Update(sm)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &parameter.SocialMediaResp{
		Id:             sm.Id,
		Name:           sm.Name,
		SocialMediaUrl: sm.SocialMediaUrl,
		UserId:         sm.UserId,
		CreatedAt:      sm.CreatedAt,
	}, http.StatusOK, nil
}

func (s SocialMediaControllerContract) Delete(req *parameter.SocialMediaReq) (int, error) {
	//TODO implement me
	sm, err := s.socialMediaRepo.GetById(req.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if sm.Id == 0 {
		return http.StatusBadRequest, errors.New(helper.CommentNotFound)
	}

	if sm.UserId != req.UserId {
		return http.StatusUnauthorized, errors.New(helper.CommentNotFound)

	}

	err = s.socialMediaRepo.Delete(sm)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func NewSocialMediaController(db *gorm.DB) SocialMediaController {
	return &SocialMediaControllerContract{
		db:              db,
		socialMediaRepo: repo.NewSocialMediaRepo(db),
	}
}
