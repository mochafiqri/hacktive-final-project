package handler

import (
	"errors"
	"finalProject/controller"
	"finalProject/helper"
	"finalProject/middleware"
	"finalProject/parameter"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SocialMediaHandler struct {
	socialMediaController controller.SocialMediaController
}

func NewSocialMediaHandler(r *gin.Engine, h *Handler) {
	handler := SocialMediaHandler{socialMediaController: controller.NewSocialMediaController(h.Db)}

	api := r.Group("socialmedias")
	api.POST("", middleware.UserAuth, handler.add)
	api.GET("", middleware.UserAuth, handler.getSocialMedia)
	api.PUT("/:id", middleware.UserAuth, handler.updateSocialMedia)
	api.DELETE("/:id", middleware.UserAuth, handler.deleteSocialMedia)
}

func (h *SocialMediaHandler) add(c *gin.Context) {
	var tmp, exist = c.Get("user_id")
	if !exist {
		helper.JSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil, nil)
		return
	}

	var req = parameter.SocialMediaReq{}
	var err = c.ShouldBindJSON(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, err.Error(), nil, err)
		return
	}

	req.UserId = tmp.(int)

	resp, status, err := h.socialMediaController.Add(&req)

	helper.JSON(c, status, http.StatusText(status), resp, err)
	return

}

func (h *SocialMediaHandler) getSocialMedia(c *gin.Context) {
	resp, status, err := h.socialMediaController.Get()
	helper.JSON(c, status, http.StatusText(status), resp, err)
	return
}

func (h *SocialMediaHandler) updateSocialMedia(c *gin.Context) {
	var userId, exist = c.Get("user_id")

	if !exist {
		helper.JSON(c, http.StatusBadRequest, helper.UserNotFound, nil, nil)
		return
	}

	var tmp = c.Param("id")
	var socialMedia, _ = strconv.Atoi(tmp)

	if socialMedia == 0 {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, errors.New(helper.SocialMediaNotFound))
		return
	}

	var req = parameter.SocialMediaReq{}
	var err = c.ShouldBind(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, err)
		return
	}
	req.Id = socialMedia
	req.UserId = userId.(int)

	resp, status, err := h.socialMediaController.Update(&req)
	helper.JSON(c, status, http.StatusText(status), resp, err)
}

func (h *SocialMediaHandler) deleteSocialMedia(c *gin.Context) {
	var userId, exist = c.Get("user_id")

	if !exist {
		helper.JSON(c, http.StatusBadRequest, helper.UserNotFound, nil, nil)
		return
	}

	var tmp = c.Param("id")
	var socialMediaId, _ = strconv.Atoi(tmp)

	if socialMediaId == 0 {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, errors.New(helper.SocialMediaNotFound))
		return
	}

	status, err := h.socialMediaController.Delete(&parameter.SocialMediaReq{
		Id:     socialMediaId,
		UserId: userId.(int),
	})

	if err != nil {
		helper.JSON(c, status, http.StatusText(status), nil, err)
		return
	}

	helper.JSON(c, status, http.StatusText(status), gin.H{
		"message": "Your comment has been successfully deleted",
	}, nil)
	return

}
