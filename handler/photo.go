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

type PhotoHandler struct {
	photoController controller.PhotoController
}

func NewPhotoHandler(r *gin.Engine, h *Handler) {
	handler := PhotoHandler{photoController: controller.NewPhotoController(h.Db)}

	api := r.Group("photos")
	api.Use(middleware.UserAuth)
	api.GET("", handler.get)
	api.POST("", handler.add)
	api.PUT("/:id", handler.update)
	api.DELETE("/:id", handler.delete)
}

func (h *PhotoHandler) get(c *gin.Context) {
	var _, exist = c.Get("user_id")
	if !exist {
		helper.JSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil, nil)
		return
	}

	resp, status, err := h.photoController.Get()
	helper.JSON(c, status, http.StatusText(status), resp, err)
	return
}

func (h *PhotoHandler) add(c *gin.Context) {
	var tmp, exist = c.Get("user_id")
	if !exist {
		helper.JSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil, nil)
		return
	}

	var req = parameter.PhotoReq{}
	var err = c.Bind(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, err)
		return
	}

	req.UserId = tmp.(int)

	resp, status, err := h.photoController.Add(&req)
	if err != nil {
		helper.JSON(c, status, http.StatusText(status), nil, err)
		return
	}

	helper.JSON(c, status, http.StatusText(status), resp, nil)
	return

}

func (h *PhotoHandler) update(c *gin.Context) {
	var tmp, exist = c.Get("user_id")
	if !exist {
		helper.JSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil, nil)
		return
	}
	var req = parameter.PhotoReq{}

	var err = c.Bind(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, err.Error(), nil, err)
		return
	}

	req.Id, _ = strconv.Atoi(c.Param("id"))
	req.UserId = tmp.(int)

	if req.Id == 0 {
		helper.JSON(c, http.StatusBadRequest, helper.PhotoNotFound, nil, errors.New(helper.PhotoNotFound))
		return
	}

	resp, status, err := h.photoController.Update(&req)
	helper.JSON(c, status, http.StatusText(status), resp, err)

}

func (h *PhotoHandler) delete(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	var tmp, exist = c.Get("user_id")
	if !exist {
		helper.JSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil, nil)
		return
	}

	if id == 0 {
		helper.JSON(c, http.StatusBadRequest, helper.PhotoNotFound, nil, errors.New(helper.PhotoNotFound))
		return
	}

	userId := tmp.(int)

	status, err := h.photoController.Delete(&parameter.PhotoReq{
		Id:     id,
		UserId: userId,
	})

	if err != nil {
		helper.JSON(c, status, http.StatusText(status), nil, err)
		return
	}

	helper.JSON(c, status, http.StatusText(status), gin.H{
		"message": "Your photo has been successfully deleted"}, err)
	return

}
