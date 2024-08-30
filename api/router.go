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

	// Initialize services
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

	// User routes
	userGroup := router.Group("/users")
	{
		userGroup.POST("/create", userHandler.Create)                          // @router /users/create [post]
		userGroup.GET("/get_profile", userHandler.GetProfile)                  // @router /users/get_profile [get]
		userGroup.PUT("/update_profile", userHandler.UpdateProfile)            // @router /users/update_profile [put]
		userGroup.PUT("/change_password", userHandler.ChangePassword)          // @router /users/change_password [put]
		userGroup.PUT("/change_profile_image", userHandler.ChangeProfileImage) // @router /users/change_profile_image [put]
		userGroup.GET("/fetch", userHandler.FetchUsers)                        // @router /users/fetch [get]
		userGroup.GET("/followers", userHandler.ListOfFollowers)               // @router /users/followers [get]
		userGroup.GET("/following", userHandler.ListOfFollowing)               // @router /users/following [get]
		userGroup.DELETE("/delete", userHandler.DeleteUser)                    // @router /users/delete [delete]
		userGroup.POST("/:id/follow", userHandler.Follow)                      // @router /users/{id}/follow [post]
		userGroup.DELETE("/:id/unfollow", userHandler.Unfollow)                // @router /users/{id}/unfollow [delete]
		userGroup.GET("/:id/followers/list", userHandler.GetUserFollowers)     // @router /users/{id}/followers/list [get]
		userGroup.GET("/:id/follows/list", userHandler.GetUserFollows)         // @router /users/{id}/follows/list [get]
		userGroup.GET("/popular", userHandler.MostPopularUser)                 // @router /users/popular [get]
	}

	// Tweet routes
	tweetGroup := router.Group("/tweets")
	{
		tweetGroup.POST("", tweetHandler.PostTweet)                 // @router /tweets [post]
		tweetGroup.GET("/:id", tweetHandler.GetTweet)               // @router /tweets/{id} [get]
		tweetGroup.PUT("/:id", tweetHandler.UpdateTweet)            // @router /tweets/{id} [put]
		tweetGroup.POST("/:id/image", tweetHandler.AddImageToTweet) // @router /tweets/{id}/image [post]
		tweetGroup.GET("/user/:userId", tweetHandler.UserTweets)    // @router /tweets/user/{userId} [get]
		tweetGroup.GET("", tweetHandler.GetAllTweets)               // @router /tweets [get]
		tweetGroup.GET("/recommend", tweetHandler.RecommendTweets)  // @router /tweets/recommend [get]
		tweetGroup.GET("/new", tweetHandler.GetNewTweets)           // @router /tweets/new [get]
	}

	// Comment routes
	commentGroup := router.Group("/comments")
	{
		commentGroup.POST("", commentHandler.PostComment)                  // @router /comments [post]
		commentGroup.GET("/:id", commentHandler.GetComment)                // @router /comments/{id} [get]
		commentGroup.PUT("/:id", commentHandler.UpdateComment)             // @router /comments/{id} [put]
		commentGroup.DELETE("/:id", commentHandler.DeleteComment)          // @router /comments/{id} [delete]
		commentGroup.GET("", commentHandler.GetAllComments)                // @router /comments [get]
		commentGroup.GET("/user/:userId", commentHandler.GetUserComments)  // @router /comments/user/{userId} [get]
		commentGroup.POST("/:id/like", commentHandler.AddLikeToComment)    // @router /comments/{id}/like [post]
		commentGroup.DELETE("/:id/like", commentHandler.DeleteLikeComment) // @router /comments/{id}/like [delete]
	}

	// Like routes
	likeGroup := router.Group("/likes")
	{
		likeGroup.POST("", likeHandler.AddLike)                                // @router /likes [post]
		likeGroup.DELETE("/:id", likeHandler.DeleteLIke)                       // @router /likes/{id} [delete]
		likeGroup.GET("/user/:userId", likeHandler.GetUserLikes)               // @router /likes/user/{userId} [get]
		likeGroup.GET("/tweet/:tweetId/count", likeHandler.GetCountTweetLikes) // @router /likes/tweet/{tweetId}/count [get]
		likeGroup.GET("/most-liked", likeHandler.MostLikedTweets)              // @router /likes/most-liked [get]
	}

	return router
}
