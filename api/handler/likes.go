package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/service"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

type LikeHandler interface {
	AddLike(c *gin.Context)
	DeleteLIke(c *gin.Context)
	GetUserLikes(c *gin.Context)
	GetCountTweetLikes(c *gin.Context)
	MostLikedTweets(c *gin.Context)
}

type HandlerDODI struct {
	TweetService pb.TweetServiceClient
	logger       *slog.Logger
}

func NewLikeHandler(tweerService service.Service, logger *slog.Logger) LikeHandler {
	tweerClient := tweerService.TweetService()
	if tweerClient == nil {
		log.Fatalf("Error creating like handler")
		return nil
	}
	return &HandlerDODI{
		TweetService: tweerClient,
		logger:       logger,
	}
}

func (h *HandlerDODI) AddLike(c *gin.Context) {

}

func (h *HandlerDODI) DeleteLIke(c *gin.Context) {

}

func (h *HandlerDODI) GetUserLikes(c *gin.Context) {

}

func (h *HandlerDODI) GetCountTweetLikes(c *gin.Context) {

}

func (h *HandlerDODI) MostLikedTweets(c *gin.Context) {

}
