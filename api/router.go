package api

import (
	"apigateway/api/handler"
	"apigateway/api/middleware"
	"apigateway/service"
	"github.com/casbin/casbin/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"

	"apigateway/pkg/config"

	_ "apigateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(cfg *config.Config, casbin *casbin.Enforcer, conn *amqp.Channel, log *slog.Logger) *gin.Engine {
	router := gin.Default()

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply middleware
	router.Use(middleware.PermissionMiddleware(casbin))
	a, err := service.NewService(cfg)
	if err != nil {
		log.Info("Error while creating service", err)
		return nil
	}
	// Create handlers with service and logger dependencies
	commentHandler := handler.NewCommentHandler(a, log)
	tweetHandler := handler.NewTweetHandler(a, log)
	userHandler := handler.NewUserHandler(a, log)
	likeHandler := handler.NewLikeHandler(a, log)

	// Comment routes
	router.POST("/comments", commentHandler.PostComment)
	router.GET("/comments/:id", commentHandler.GetComment)
	router.PUT("/comments/:id", commentHandler.UpdateComment)
	router.DELETE("/comments/:id", commentHandler.DeleteComment)
	router.GET("/comments", commentHandler.GetAllComments)
	router.GET("/comments/user/:userId", commentHandler.GetUserComments)
	router.POST("/comments/:id/like", commentHandler.AddLikeToComment)
	router.DELETE("/comments/:id/like", commentHandler.DeleteLikeComment)

	// Tweet routes
	router.POST("/tweets", tweetHandler.PostTweet)
	router.GET("/tweets/:id", tweetHandler.GetTweet)
	router.PUT("/tweets/:id", tweetHandler.UpdateTweet)
	router.POST("/tweets/:id/image", tweetHandler.AddImageToTweet)
	router.GET("/tweets/user/:userId", tweetHandler.UserTweets)
	router.GET("/tweets", tweetHandler.GetAllTweets)
	router.GET("/tweets/recommend", tweetHandler.RecommendTweets)
	router.GET("/tweets/new", tweetHandler.GetNewTweets)

	// Like routes
	router.POST("/likes", likeHandler.AddLike)
	router.DELETE("/likes/:id", likeHandler.DeleteLIke)
	router.GET("/likes/user/:userId", likeHandler.GetUserLikes)
	router.GET("/likes/tweet/:tweetId/count", likeHandler.GetCountTweetLikes)
	router.GET("/likes/most-liked", likeHandler.MostLikedTweets)

	// User routes
	router.POST("/users", userHandler.Create)
	router.GET("/users/:id", userHandler.GetProfile)
	router.PUT("/users/:id", userHandler.UpdateProfile)
	router.PUT("/users/:id/password", userHandler.ChangePassword)
	router.PUT("/users/:id/profile-image", userHandler.ChangeProfileImage)
	router.GET("/users", userHandler.FetchUsers)
	router.GET("/users/:id/following", userHandler.ListOfFollowing)
	router.GET("/users/:id/followers", userHandler.ListOfFollowers)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	// User follow/unfollow routes
	router.POST("/users/:id/follow", userHandler.Follow)
	router.DELETE("/users/:id/unfollow", userHandler.Unfollow)
	router.GET("/users/:id/followers/list", userHandler.GetUserFollowers)
	router.GET("/users/:id/follows/list", userHandler.GetUserFollows)
	router.GET("/users/popular", userHandler.MostPopularUser)

	return router
}
