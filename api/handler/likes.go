package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/pkg/models"
	"apigateway/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
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
	var like models.LikeReq
	if err := c.ShouldBindJSON(&like); err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := pb.LikeReq{
		UserId:  c.Value("user_id").(string),
		TweetId: like.TweetID,
	}

	req, err := h.TweetService.AddLike(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *HandlerDODI) DeleteLIke(c *gin.Context) {
	var like models.LikeReq
	if err := c.ShouldBindJSON(&like); err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := pb.LikeReq{
		UserId:  c.Value("user_id").(string),
		TweetId: like.TweetID,
	}

	req, err := h.TweetService.AddLike(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *HandlerDODI) GetUserLikes(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}
	req, err := h.TweetService.GetUserLikes(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *HandlerDODI) GetCountTweetLikes(c *gin.Context) {
	tweetid := c.Param("tweet_id")

	res := pb.TweetId{
		Id: tweetid,
	}

	req, err := h.TweetService.GetCountTweetLikes(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *HandlerDODI) MostLikedTweets(c *gin.Context) {
	res := pb.Void{}

	req, err := h.TweetService.MostLikedTweets(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}
