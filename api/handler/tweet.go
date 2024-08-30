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

// PostTweet godoc
// @Summary PostTweet Tweets
// @Description sign in tweet
// @Tags Tweet
// @Accept json
// @Produce json
// @Param PostTweet body models.Tweet true "post tweet"
// @Success 200 {object} models.TweetResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/add [post]
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

// UpdateTweet godoc
// @Summary UpdateTweet Tweets
// @Description sign in tweet
// @Tags Tweet
// @Accept json
// @Produce json
// @Param UpdateTweet body models.UpdateATweet true "put like"
// @Success 200 {object} models.TweetResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/add [put]
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

// AddImageToTweet godoc
// @Summary AddImageToTweet Tweets
// @Description sign in tweet
// @Tags Tweet
// @Accept json
// @Produce json
// @Param AddImageToTweet body models.Url true "get like"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/get [get]
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

// UserTweets godoc
// @Summary UserTweets Tweets
// @Description sign in tweet
// @Tags Tweet
// @Accept json
// @Produce json
// @Param UserTweets body models.Id true "get like"
// @Success 200 {object} models.Tweets
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/user [get]
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

// GetTweet godoc
// @Summary GetTweet Tweets
// @Description sign in tweet
// @Tags Tweet
// @Accept json
// @Produce json
// @Param GetTweet body models.TweetId true "get like"
// @Success 200 {object} models.TweetResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/get_tt [get]
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

// GetAllTweets godoc
// @Summary GetAllTweets Tweets
// @Description sign in tweet
// @Tags Tweet
// @Accept json
// @Produce json
// @Param GetAllTweets body models.TweetFilter true "get like"
// @Success 200 {object} models.Tweets
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/get_all [get]
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

// RecommendTweets godoc
// @Summary RecommendTweets Tweets
// @Description sign in tweet
// @Tags Tweet
// @Accept json
// @Produce json
// @Param RecommendTweets body models.Id true "get like"
// @Success 200 {object} models.Tweets
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/recommend [get]
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

// GetNewTweets godoc
// @Summary GetNewTweets Tweets
// @Description sign in tweet
// @Tags Tweet
// @Accept json
// @Produce json
// @Param GetNewTweets body models.Id true "get like"
// @Success 200 {object} models.Tweets
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/get_new [get]
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
