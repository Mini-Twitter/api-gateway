package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/pkg/models"
	"apigateway/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"log/slog"
	"net/http"
)

type CommentHandler interface {
	PostComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	GetComment(c *gin.Context)
	GetAllComments(c *gin.Context)
	GetUserComments(c *gin.Context)
	AddLikeToComment(c *gin.Context)
	DeleteLikeComment(c *gin.Context)
}

type commentHandler struct {
	CommentService pb.TweetServiceClient
	logger         *slog.Logger
}

func NewCommentHandler(commentService service.Service, logger *slog.Logger) CommentHandler {
	commentClient := commentService.TweetService()
	if commentClient == nil {
		log.Fatalf("Failed to create comment handler")
		return nil
	}
	return &commentHandler{
		CommentService: commentClient,
		logger:         logger,
	}
}

// PostComment godoc
// @Summary PostComment Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param PostComment body models.Comment true "post comment"
// @Success 200 {object} models.CommentRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/post [post]
func (h *commentHandler) PostComment(c *gin.Context) {
	var tweet models.Comment
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := pb.Comment{
		Id:        uuid.NewString(),
		UserId:    c.MustGet("user_id").(string),
		TweetId:   tweet.TweetID,
		Content:   tweet.Content,
		LikeCount: tweet.LikeCount,
	}

	req, err := h.CommentService.PostComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// UpdateComment godoc
// @Summary UpdateComment Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param UpdateComment body models.UpdateAComment true "put comment"
// @Success 200 {object} models.CommentRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/update [put]
func (h *commentHandler) UpdateComment(c *gin.Context) {
	var tweet models.UpdateAComment
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.UpdateAComment{
		Id:      c.MustGet("user_id").(string),
		Content: tweet.Content,
	}
	req, err := h.CommentService.UpdateComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while updating comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// DeleteComment godoc
// @Summary DeleteComment Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param DeleteComment body models.CommentId true "delete comment"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/delete [delete]
func (h *commentHandler) DeleteComment(c *gin.Context) {
	id := c.Param("id")

	res := pb.CommentId{
		Id: id,
	}

	req, err := h.CommentService.DeleteComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetComment godoc
// @Summary GetComment Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param GetComment body models.CommentId true "get comment"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/get [get]
func (h *commentHandler) GetComment(c *gin.Context) {
	id := c.Param("id")

	res := pb.CommentId{
		Id: id,
	}

	req, err := h.CommentService.DeleteComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetAllComments godoc
// @Summary GetAllComments Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce jsonUserId
// @Param GetAllComments body models.CommentFilter true "get comment"
// @Success 200 {object} models.Comments
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/get_all [get]
func (h *commentHandler) GetAllComments(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	tweetId := c.Param("tweet_id")

	res := pb.CommentFilter{
		UserId:  userId,
		TweetId: tweetId,
	}
	req, err := h.CommentService.GetAllComments(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting comments", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetUserComments godoc
// @Summary GetUserComments Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param GetUserComments body models.Id true "get comment"
// @Success 200 {object} models.Comments
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/get_user [get]
func (h *commentHandler) GetUserComments(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}
	req, err := h.CommentService.GetUserComments(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting comments", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// AddLikeToComment godoc
// @Summary AddLikeToComment Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param AddLikeToComment body models.CommentLikeReq true "get comment"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/get_add [get]
func (h *commentHandler) AddLikeToComment(c *gin.Context) {
	id := c.Param("id")
	res := pb.CommentLikeReq{
		CommentId: id,
	}
	req, err := h.CommentService.AddLikeToComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while adding like to comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// DeleteLikeComment godoc
// @Summary DeleteLikeComment Comments
// @Description sign in comment
// @Tags Tweet
// @Accept json
// @Produce json
// @Param DeleteLikeComment body models.CommentLikeReq true "get comment"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/delete [delete]
func (h *commentHandler) DeleteLikeComment(c *gin.Context) {
	id := c.Param("id")
	res := pb.CommentLikeReq{
		CommentId: id,
	}
	req, err := h.CommentService.DeleteLikeComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting like comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}
