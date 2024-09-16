package handler

import (
	"github.com/Sabyradinov/go-dating-matchmaking/common"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/logger"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Matching struct {
	log logger.AppLogger
	srv service.IMatching
}

func NewMatching(opt *Options) *Matching {
	return &Matching{log: opt.Logger, srv: opt.Services.Matching}
}

// GetPotentialMatches method to get potential matches
// @Summary method to get potential matches
// @Description method to recommend potential matches for a user based on certain criteria, such as preferences, location, and mutual interests
// @Tags Matching
// @Accept  json
// @Produce  json
// @Param userId query string true "userId"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} model.UserResponse "response body"
// @Failure 400,404 {object} common.BaseResponse{} "error body"
// @Router /match/recommendations [get]
func (h Matching) GetPotentialMatches(ctx *gin.Context) {
	userId, ok := ctx.GetQuery("userId")
	if !ok {
		h.log.ErrorWithCode(ctx, "GetPotentialMatches", -911, "validation error", nil)
		return
	}
	pageStr := ctx.Query("page")
	sizeStr := ctx.Query("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 10
	}

	res, err := h.srv.GetPotentialMatches(ctx, userId, page, size)
	if err != nil {
		h.log.ErrorWithCode(ctx, "GetPotentialMatches", -911, err.Error(), nil)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetUserById method to get user by id
// @Summary method to get user by id
// @Description method to get user by id
// @Tags Matching
// @Accept  json
// @Produce  json
// @Param userId query string true "userId"
// @Success 200 {object} common.BaseResponse{} "response body"
// @Failure 400,404 {object} common.BaseResponse{} "error body"
// @Router /matching/curren-user/:id [get]
func (h Matching) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if len(userId) == 0 {
		h.log.ErrorWithCode(ctx, "GetPotentialMatches", -911, "validation error", nil)
		return
	}

	res, err := h.srv.GetUserById(ctx, userId)
	if err != nil {
		h.log.ErrorWithCode(ctx, "GetUserById", -911, err.Error(), nil)
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{
		Code:    0,
		Message: "success",
		Data:    res,
	})
}
