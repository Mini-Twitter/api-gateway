package handler

import (
	pb "apigateway/genproto/user"
	"apigateway/service"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

type UserHandler interface {
	Create(c *gin.Context)
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	ChangePassword(c *gin.Context)
	ChangeProfileImage(c *gin.Context)
	FetchUsers(c *gin.Context)
	ListOfFollowing(c *gin.Context)
	ListOfFollowers(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	userService pb.UserServiceClient
	logger      *slog.Logger
}

func NewUserHandler(userService service.Service, logger *slog.Logger) UserHandler {
	userClient := userService.UserService()
	if userClient == nil {
		log.Fatalf("failed to create user client")
		return nil
	}
	return &userHandler{
		userService: userClient,
		logger:      logger,
	}
}

func (h *userHandler) Create(c *gin.Context) {

}

func (h *userHandler) GetProfile(c *gin.Context) {

}

func (h *userHandler) UpdateProfile(c *gin.Context) {

}

func (h *userHandler) ChangePassword(c *gin.Context) {

}

func (h *userHandler) ChangeProfileImage(c *gin.Context) {

}

func (h *userHandler) FetchUsers(c *gin.Context) {

}

func (h *userHandler) ListOfFollowing(c *gin.Context) {

}

func (h *userHandler) ListOfFollowers(c *gin.Context) {

}

func (h *userHandler) DeleteUser(c *gin.Context) {

}
