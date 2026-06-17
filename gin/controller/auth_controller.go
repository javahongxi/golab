package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/javahongxi/golab/gin/model"
	"github.com/javahongxi/golab/gin/response"
	"github.com/javahongxi/golab/gin/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthController struct {
	repo model.UserRepository
}

func NewAuthController(repo model.UserRepository) *AuthController {
	return &AuthController{repo: repo}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Nickname string `json:"nickname" binding:"max=20"`
	Gender   int32  `json:"gender" binding:"min=0,max=2"`
	Age      int32  `json:"age" binding:"min=0,max=150"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest
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
		zap.L().Error("failed to register user", zap.Error(err))
		response.ServerError(ctx, "failed to register")
		return
	}

	response.Success(ctx, user)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	user, err := c.repo.FindByUsername(req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.ErrorWithCode(ctx, 401, "invalid username or password")
			return
		}
		zap.L().Error("failed to find user", zap.Error(err))
		response.ServerError(ctx, "login failed")
		return
	}

	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		zap.L().Error("failed to generate token", zap.Error(err))
		response.ServerError(ctx, "login failed")
		return
	}

	response.Success(ctx, LoginResponse{
		User:  user,
		Token: token,
	})
}
