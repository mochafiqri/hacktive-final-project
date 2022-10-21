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

type CommentHandler struct {
	commentController controller.CommentController
}

func NewCommentHandler(r *gin.Engine, h *Handler) {
	var handler = CommentHandler{commentController: controller.NewCommentController(h.Db)}

	api := r.Group("comments")
	api.POST("", middleware.UserAuth, handler.addComment)
	api.GET("", middleware.UserAuth, handler.getComment)
	api.PUT("/:id", middleware.UserAuth, handler.updateComment)
	api.DELETE("/:id", middleware.UserAuth, handler.deleteComment)

}

func (h *CommentHandler) addComment(c *gin.Context) {
	var tmp, exist = c.Get("user_id")
	if !exist {
		helper.JSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil, nil)
		return
	}

	var req = parameter.CommentReq{}
	var err = c.ShouldBindJSON(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, err.Error(), nil, err)
		return
	}

	req.UserId = tmp.(int)

	resp, status, err := h.commentController.AddComment(&req)

	helper.JSON(c, status, http.StatusText(status), resp, err)
	return

}

func (h *CommentHandler) getComment(c *gin.Context) {
	resp, status, err := h.commentController.Get()
	helper.JSON(c, status, http.StatusText(status), resp, err)
	return
}

func (h *CommentHandler) updateComment(c *gin.Context) {
	var userId, exist = c.Get("user_id")

	if !exist {
		helper.JSON(c, http.StatusBadRequest, helper.UserNotFound, nil, nil)
		return
	}

	var tmp = c.Param("id")
	var commentId, _ = strconv.Atoi(tmp)

	if commentId == 0 {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, errors.New(helper.CommentNotFound))
		return
	}

	var req = parameter.CommentUpdate{}
	var err = c.ShouldBind(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, err)
		return
	}

	resp, status, err := h.commentController.Update(&parameter.CommentReq{Id: commentId, Message: req.Message, UserId: userId.(int)})
	helper.JSON(c, status, http.StatusText(status), resp, err)
}

func (h *CommentHandler) deleteComment(c *gin.Context) {
	var userId, exist = c.Get("user_id")

	if !exist {
		helper.JSON(c, http.StatusBadRequest, helper.UserNotFound, nil, nil)
		return
	}

	var tmp = c.Param("id")
	var commentId, _ = strconv.Atoi(tmp)

	if commentId == 0 {
		helper.JSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, errors.New(helper.CommentNotFound))
		return
	}

	status, err := h.commentController.Delete(&parameter.CommentReq{Id: commentId, UserId: userId.(int)})
	if err != nil {
		helper.JSON(c, status, http.StatusText(status), nil, err)
		return
	}
	helper.JSON(c, status, http.StatusText(status), gin.H{
		"message": "Your comment has been successfully deleted",
	}, nil)
	return

}
