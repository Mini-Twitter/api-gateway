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
