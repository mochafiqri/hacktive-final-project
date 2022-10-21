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

type UserHandlerContract struct {
	controller.UserController
}

func NewUserHandler(r *gin.Engine, h *Handler) {
	handler := UserHandlerContract{
		UserController: controller.NewUserController(h.Db),
	}
	api := r.Group("users")
	api.POST("/register", handler.register)
	api.POST("/login", handler.login)
	api.PUT("/:id", middleware.UserAuth, handler.update)
	api.DELETE("", middleware.UserAuth, handler.delete)
}

func (h *UserHandlerContract) register(c *gin.Context) {
	var req = parameter.RegisterReq{}
	var err = c.ShouldBindJSON(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, err)
		return
	}

	if !helper.ValidEmail(req.Email) { // Check Email Valid
		helper.JSON(c, http.StatusBadRequest, helper.EmailNotValid, nil, errors.New(helper.EmailNotValid))
		return
	}

	if req.Age <= 8 {
		helper.JSON(c, http.StatusBadRequest, helper.AgeMinimum, nil, errors.New(helper.AgeMinimum))
		return
	}

	if len(req.Password) < 6 {
		helper.JSON(c, http.StatusBadRequest, helper.PasswordInvalid, nil, errors.New(helper.PasswordInvalid))
		return
	}

	resp, status, err := h.UserController.Register(&req)
	if err != nil {
		helper.JSON(c, status, "failed", nil, err)
		return
	}

	if !helper.ValidEmail(req.Email) { // Check Email Valid
		helper.JSON(c, http.StatusBadRequest, helper.EmailNotValid, nil, errors.New(helper.EmailNotValid))
		return
	}

	helper.JSON(c, status, "success register", resp, nil)
	return
}

func (h *UserHandlerContract) login(c *gin.Context) {
	var req = parameter.LoginReq{}

	var err = c.ShouldBindJSON(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, err)
		return
	}

	if len(req.Password) < 6 {
		helper.JSON(c, http.StatusBadRequest, helper.UserNotFound, nil, nil)
		return
	}
	resp, status, err := h.UserController.Login(&req)

	helper.JSON(c, status, http.StatusText(status), resp, err)
	return

}

func (h *UserHandlerContract) update(c *gin.Context) {
	var req = parameter.UpdateReq{}
	var tmp = c.Param("id")
	var err = c.ShouldBind(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, err)
		return
	}
	var id, _ = strconv.Atoi(tmp)
	if id == 0 {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, errors.New(helper.UserNotFound))
		return
	}

	req.Id = id
	resp, status, err := h.UserController.Update(&req)

	helper.JSON(c, status, http.StatusText(status), resp, err)
	return

}

func (h *UserHandlerContract) delete(c *gin.Context) {
	var tmp, exist = c.Get("user_id")

	if !exist {
		helper.JSON(c, http.StatusBadRequest, helper.UserNotFound, nil, nil)
		return
	}

	status, err := h.UserController.Delete(tmp.(int))
	if err != nil {
		helper.JSON(c, status, err.Error(), nil, err)
		return
	}

	if status != http.StatusOK {
		helper.JSON(c, status, http.StatusText(status), nil, nil)
		return
	}

	helper.JSON(c, status, http.StatusText(status), gin.H{
		"message": "Your account has been successfully deleted",
	}, nil)
	return

}
