package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/pkg/models"
	"apigateway/service"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"
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
	var tweet models.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := pb.Tweet{
		UserId:    c.MustGet("user_id").(string),
		Hashtag:   tweet.Hashtag,
		Title:     tweet.Title,
		Content:   tweet.Content,
		ImageUrl:  tweet.ImageURL,
		CreatedAt: time.Now().String(),
		LikeCount: tweet.LikeCount,
	}

	req, err := h.tweetService.PostTweet(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *tweetHandler) UpdateTweet(c *gin.Context) {
	var tweet models.UpdateATweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.UpdateATweet{
		Id:      c.MustGet("id").(string),
		Hashtag: tweet.Hashtag,
		Title:   tweet.Title,
		Content: tweet.Content,
	}
	req, err := h.tweetService.UpdateTweet(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while updating tweet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *tweetHandler) AddImageToTweet(c *gin.Context) {
	var tweet models.Url
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.Url{
		TweetId: tweet.TweetID,
		Url:     tweet.URL,
	}

	req, err := h.tweetService.AddImageToTweet(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while adding image to tweet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *tweetHandler) UserTweets(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}

	req, err := h.tweetService.UserTweets(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting user tweets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *tweetHandler) GetTweet(c *gin.Context) {
	id := c.Param("id")

	res := pb.TweetId{
		Id: id,
	}

	req, err := h.tweetService.GetTweet(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting tweet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *tweetHandler) GetAllTweets(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		offsets = 1
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	hashtag := c.Query("hashtag")
	title := c.Query("title")

	res := pb.TweetFilter{
		Limit:   int64(limits),
		Offset:  int64(offsets),
		Hashtag: hashtag,
		Title:   title,
	}

	req, err := h.tweetService.GetAllTweets(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting tweets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *tweetHandler) RecommendTweets(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}
	req, err := h.tweetService.RecommendTweets(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting tweets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *tweetHandler) GetNewTweets(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}
	req, err := h.tweetService.RecommendTweets(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting tweets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}
