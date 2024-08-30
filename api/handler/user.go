package handler

import (
	pb "apigateway/genproto/user"
	"apigateway/pkg/models"
	"apigateway/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
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

	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetUserFollows(c *gin.Context)
	MostPopularUser(c *gin.Context)
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
	var user models.CreateRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := pb.CreateRequest{
		Email:       user.Email,
		Password:    user.Password,
		Phone:       user.Phone,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		Nationality: user.Nationality,
		Bio:         user.Bio,
	}

	res, err := h.userService.Create(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while creating user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *userHandler) GetProfile(c *gin.Context) {
	req := pb.Id{
		UserId: c.MustGet("UserId").(string),
	}

	res, err := h.userService.GetProfile(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while getting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) UpdateProfile(c *gin.Context) {
	var user models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.UpdateProfileRequest{
		UserId:       c.MustGet("UserId").(string),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PhoneNumber:  user.PhoneNumber,
		Username:     user.Username,
		Nationality:  user.Nationality,
		Bio:          user.Bio,
		ProfileImage: user.ProfileImage,
	}

	req, err := h.userService.UpdateProfile(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while updating user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) ChangePassword(c *gin.Context) {
	var user models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.ChangePasswordRequest{
		UserId:          c.MustGet("UserId").(string),
		CurrentPassword: user.CurrentPassword,
		NewPassword:     user.NewPassword,
	}

	req, err := h.userService.ChangePassword(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while changing user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) ChangeProfileImage(c *gin.Context) {
	var user models.URL
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.URL{
		UserId: "",
		Url:    user.URL,
	}
	req, err := h.userService.ChangeProfileImage(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while changing user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) FetchUsers(c *gin.Context) {
	role := c.Query("role")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	name := c.Query("name")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	res := pb.Filter{
		Role:      role,
		Page:      int32(page),
		Limit:     int32(limit),
		FirstName: name,
	}

	req, err := h.userService.FetchUsers(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while fetching users", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)

}

func (h *userHandler) ListOfFollowing(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("UserId").(string),
	}
	req, err := h.userService.ListOfFollowing(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while listing following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) ListOfFollowers(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("UserId").(string),
	}
	req, err := h.userService.ListOfFollowers(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while listing followers", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("UserId").(string),
	}
	req, err := h.userService.DeleteUser(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) Follow(c *gin.Context) {
	var user models.FollowReq
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.FollowReq{
		FollowingId: user.FollowingID,
		FollowerId:  user.FollowerID,
	}

	req, err := h.userService.Follow(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) Unfollow(c *gin.Context) {
	var user models.FollowReq
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.FollowReq{
		FollowingId: user.FollowingID,
		FollowerId:  user.FollowerID,
	}

	req, err := h.userService.Unfollow(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) GetUserFollowers(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("UserId").(string),
	}
	req, err := h.userService.GetUserFollowers(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) GetUserFollows(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("UserId").(string),
	}
	req, err := h.userService.GetUserFollows(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *userHandler) MostPopularUser(c *gin.Context) {
	res := pb.Void{}
	req, err := h.userService.MostPopularUser(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while most popular user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}
