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

type CommentController interface {
	AddComment(req *parameter.CommentReq) (*parameter.CommentResp, int, error)
	Get() ([]parameter.CommentList, int, error)
	Update(req *parameter.CommentReq) (*parameter.CommentResp, int, error)
	Delete(req *parameter.CommentReq) (int, error)
}

type CommentControllerContract struct {
	db          *gorm.DB
	commentRepo repo.CommentRepo
	photoRepo   repo.PhotoRepo
}

func (c CommentControllerContract) Delete(req *parameter.CommentReq) (int, error) {
	//TODO implement me
	comment, err := c.commentRepo.GetById(req.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if comment.Id == 0 {
		return http.StatusBadRequest, errors.New(helper.CommentNotFound)
	}

	if comment.UserId != req.UserId {
		return http.StatusUnauthorized, errors.New(helper.CommentNotFound)

	}

	err = c.commentRepo.Delete(comment)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (c CommentControllerContract) Get() ([]parameter.CommentList, int, error) {
	//TODO implement me
	var resp = []parameter.CommentList{}
	comments, err := c.commentRepo.Get()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(comments) == 0 {
		return nil, http.StatusNotFound, errors.New(helper.CommentNotFound)
	}

	for _, v := range comments {
		var tmp = parameter.CommentList{
			Id:        v.Id,
			Message:   v.Message,
			PhotoId:   v.PhotoId,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: parameter.UserResp{
				Id:       v.User.Id,
				Username: v.User.Username,
				Email:    v.User.Email,
			},
			Photo: parameter.PhotoResp{
				Id:       v.Photo.Id,
				Title:    v.Photo.Title,
				Caption:  v.Photo.Caption,
				PhotoUrl: v.Photo.PhotoUrl,
				UserId:   v.Photo.UserId,
				Created:  v.Photo.CreatedAt,
			},
		}

		resp = append(resp, tmp)
	}
	return resp, http.StatusOK, nil
}

func (c CommentControllerContract) Update(req *parameter.CommentReq) (*parameter.CommentResp, int, error) {
	//TODO implement me
	comment, err := c.commentRepo.GetById(req.Id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if comment.Id == 0 {
		return nil, http.StatusBadRequest, errors.New(helper.CommentNotFound)
	}

	if comment.UserId != req.UserId {
		return nil, http.StatusUnauthorized, errors.New(helper.CommentNotFound)

	}
	comment.Message = req.Message

	err = c.commentRepo.Update(comment)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &parameter.CommentResp{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: comment.CreatedAt,
	}, http.StatusOK, nil
}

func (c CommentControllerContract) AddComment(req *parameter.CommentReq) (*parameter.CommentResp, int, error) {
	//TODO implement me
	photo, err := c.photoRepo.Get(req.PhotoId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if photo.Id == 0 {
		return nil, http.StatusBadRequest, errors.New(helper.PhotoNotFound)
	}

	var comment = model.Comment{
		UserId:    req.UserId,
		PhotoId:   req.PhotoId,
		Message:   req.Message,
		CreatedAt: time.Now(),
	}
	err = c.commentRepo.Create(&comment)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var resp = parameter.CommentResp{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: comment.CreatedAt,
	}
	return &resp, http.StatusOK, nil
}

func NewCommentController(db *gorm.DB) CommentController {
	return &CommentControllerContract{
		db:          db,
		commentRepo: repo.NewCommentRepo(db),
		photoRepo:   repo.NewPhotoRepo(db),
	}
}
