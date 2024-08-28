package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/service"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

type SubscribeHandler interface {
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetUserFollows(c *gin.Context)
	MostPopularUser(c *gin.Context)
}

type subscribeHandler struct {
	subscribeService pb.TweetServiceClient
	logger           *slog.Logger
}

func NewSubscribeHandler(subscribeService service.Service, logger *slog.Logger) SubscribeHandler {
	subscribeClinet := subscribeService.TweetService()
	if subscribeClinet == nil {
		log.Fatalf("Failed to create subscribe handler")
		return nil
	}
	return &subscribeHandler{
		subscribeService: subscribeClinet,
		logger:           logger,
	}
}

func (h *subscribeHandler) Follow(c *gin.Context) {

}

func (h *subscribeHandler) Unfollow(c *gin.Context) {

}

func (h *subscribeHandler) GetUserFollowers(c *gin.Context) {

}

func (h *subscribeHandler) GetUserFollows(c *gin.Context) {

}

func (h *subscribeHandler) MostPopularUser(c *gin.Context) {

}
