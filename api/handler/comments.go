package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/service"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
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

}

func (h *commentHandler) UpdateComment(c *gin.Context) {

}

func (h *commentHandler) DeleteComment(c *gin.Context) {

}

func (h *commentHandler) GetComment(c *gin.Context) {

}

func (h *commentHandler) GetAllComments(c *gin.Context) {

}

func (h *commentHandler) GetUserComments(c *gin.Context) {

}

func (h *commentHandler) AddLikeToComment(c *gin.Context) {

}

func (h *commentHandler) DeleteLikeComment(c *gin.Context) {

}
