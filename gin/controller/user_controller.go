package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/javahongxi/golab/gin/model"
	"github.com/javahongxi/golab/gin/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserController struct {
	repo model.UserRepository
}

func NewUserController(repo model.UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	_, err := c.repo.FindByUsername(req.Username)
	if err == nil {
		response.ErrorWithCode(ctx, 409, "username already exists")
		return
	}

	user := &model.User{
		Username: req.Username,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		Age:      req.Age,
	}

	if err := c.repo.Create(user); err != nil {
		zap.L().Error("failed to create user", zap.Error(err))
		response.ServerError(ctx, "failed to create user")
		return
	}

	response.Success(ctx, user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "invalid user ID")
		return
	}

	user, err := c.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(ctx)
			return
		}
		zap.L().Error("failed to get user", zap.Error(err))
		response.ServerError(ctx, "failed to get user")
		return
	}

	response.Success(ctx, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "invalid user ID")
		return
	}

	var req model.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	user, err := c.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(ctx)
			return
		}
		zap.L().Error("failed to get user", zap.Error(err))
		response.ServerError(ctx, "failed to get user")
		return
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Age > 0 {
		user.Age = req.Age
	}

	if err := c.repo.Update(user); err != nil {
		zap.L().Error("failed to update user", zap.Error(err))
		response.ServerError(ctx, "failed to update user")
		return
	}

	response.Success(ctx, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(ctx, "invalid user ID")
		return
	}

	user, err := c.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(ctx)
			return
		}
		zap.L().Error("failed to get user", zap.Error(err))
		response.ServerError(ctx, "failed to get user")
		return
	}

	if err := c.repo.Delete(id); err != nil {
		zap.L().Error("failed to delete user", zap.Error(err))
		response.ServerError(ctx, "failed to delete user")
		return
	}

	response.SuccessWithMsg(ctx, "user deleted", user)
}

func (c *UserController) ListUsers(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	users, total, err := c.repo.List(page, limit)
	if err != nil {
		zap.L().Error("failed to list users", zap.Error(err))
		response.ServerError(ctx, "failed to list users")
		return
	}

	response.Success(ctx, gin.H{
		"data":  users,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}