package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/service"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

type TweetHandler interface {
	PostTweet(c *gin.Context)
	UpdateTweet(c *gin.Context)
	AddImageToTweet(c *gin.Context)
	UserTweets(c *gin.Context)
	GetTweet(c *gin.Context)
	GetAllTweets(c *gin.Context)
	RecommendTweets(c *gin.Context)
	GetNewTweets(c *gin.Context)
}

type tweetHandler struct {
	tweetService pb.TweetServiceClient
	logger       *slog.Logger
}

func NewTweetHandler(tweetService service.Service, logger *slog.Logger) TweetHandler {
	tweetClient := tweetService.TweetService()
	if tweetClient == nil {
		log.Fatalf("tweet client is nil")
		return nil
	}
	return &tweetHandler{
		tweetService: tweetClient,
		logger:       logger,
	}
}

func (h *tweetHandler) PostTweet(c *gin.Context) {

}

func (h *tweetHandler) UpdateTweet(c *gin.Context) {

}

func (h *tweetHandler) AddImageToTweet(c *gin.Context) {

}

func (h *tweetHandler) UserTweets(c *gin.Context) {

}

func (h *tweetHandler) GetTweet(c *gin.Context) {

}

func (h *tweetHandler) GetAllTweets(c *gin.Context) {

}

func (h *tweetHandler) RecommendTweets(c *gin.Context) {

}

func (h *tweetHandler) GetNewTweets(c *gin.Context) {

}
