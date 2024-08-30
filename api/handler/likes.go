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

// AddLike godoc
// @Summary AddLike Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param AddLike body models.LikeReq true "post like"
// @Success 200 {object} models.LikeRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/add [post]
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

// DeleteLIke godoc
// @Summary DeleteLIke Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param DeleteLIke body models.LikeReq true "delete like"
// @Success 200 {object} models.LikeRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/delete [delete]
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

// GetUserLikes godoc
// @Summary GetUserLikes Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param GetUserLikes body models.Id true "get like"
// @Success 200 {object} models.TweetTitles
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/get_user [get]
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

// GetCountTweetLikes godoc
// @Summary GetCountTweetLikes Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param GetCountTweetLikes body models.TweetId true "get like"
// @Success 200 {object} models.Count
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/get_count [get]
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

// MostLikedTweets godoc
// @Summary MostLikedTweets Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param MostLikedTweets body models.Void true "get like"
// @Success 200 {object} models.Tweet
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/get_most [get]
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
