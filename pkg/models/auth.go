package models

// Void message uchun struktura
type Void struct{}

// Id message uchun struktura
type Id struct {
	UserID string `json:"user_id" db:"user_id"`
}

// CreateRequest message uchun struktura
type CreateRequest struct {
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	Phone       string `json:"phone" db:"phone"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Username    string `json:"username" db:"username"`
	Nationality string `json:"nationality" db:"nationality"`
	Bio         string `json:"bio" db:"bio"`
}

// UserResponse message uchun struktura
type UserResponse struct {
	Id          string `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Username    string `json:"username" db:"username"`
	Nationality string `json:"nationality" db:"nationality"`
	Bio         string `json:"bio" db:"bio"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

// LoginRequest message uchun struktura
type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

// LoginResponse message uchun struktura
type LoginResponse struct {
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	UserID       string `json:"user_id" db:"user_id"`
}

// GetProfileResponse message uchun struktura
type GetProfileResponse struct {
	UserID         string `json:"user_id" db:"user_id"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	Email          string `json:"email" db:"email"`
	PhoneNumber    string `json:"phone_number" db:"phone_number"`
	Username       string `json:"username" db:"username"`
	Nationality    string `json:"nationality" db:"nationality"`
	Bio            string `json:"bio" db:"bio"`
	ProfileImage   string `json:"profile_image" db:"profile_image"`
	FollowersCount int32  `json:"followers_count" db:"followers_count"`
	FollowingCount int32  `json:"following_count" db:"following_count"`
	PostsCount     int32  `json:"posts_count" db:"posts_count"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
}

// UpdateProfileRequest message uchun struktura
type UpdateProfileRequest struct {
	UserID       string `json:"user_id" db:"user_id"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
	Username     string `json:"username" db:"username"`
	Nationality  string `json:"nationality" db:"nationality"`
	Bio          string `json:"bio" db:"bio"`
	ProfileImage string `json:"profile_image" db:"profile_image"`
	Phone        string `json:"phone" db:"phone"`
}

// Filter message uchun struktura
type Filter struct {
	Role      string `json:"role" db:"role"`
	Page      int32  `json:"page" db:"page"`
	Limit     int32  `json:"limit" db:"limit"`
	FirstName string `json:"first_name" db:"first_name"`
}

// UserResponses message uchun struktura
type UserResponses struct {
	Users []UserResponse `json:"users" db:"users"`
}

// ChangePasswordRequest message uchun struktura
type ChangePasswordRequest struct {
	UserID          string `json:"userId" db:"user_id"`
	CurrentPassword string `json:"current_password" db:"current_password"`
	NewPassword     string `json:"new_password" db:"new_password"`
}

// ChangePasswordResponse message uchun struktura
type ChangePasswordResponse struct {
	Message string `json:"message" db:"message"`
}

// URL message uchun struktura
type URL struct {
	URL    string `json:"url" db:"url"`
	UserID string `json:"user_id" db:"user_id"`
}

// Ids message uchun struktura
type Ids struct {
	FollowerID  string `json:"follower_id" db:"follower_id"`
	FollowingID string `json:"following_id" db:"following_id"`
}

// Followings message uchun struktura
type Followings struct {
	Ids []Ids `json:"ids" db:"ids"`
}

// Follower message uchun struktura
type Follower struct {
	UserID   string `json:"user_id" db:"user_id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
}

// Followers message uchun struktura
type Followers struct {
	Followers []Follower `json:"followers" db:"followers"`
}

// DFollowRes corresponds to the DFollowRes message
type DFollowRes struct {
	FollowerID   string `json:"follower_id" db:"follower_id"`
	FollowingID  string `json:"following_id" db:"following_id"`
	UnfollowedAt string `json:"unfollowed_at" db:"unfollowed_at"`
}

// Count corresponds to the Count message
type Count struct {
	Description string `json:"description" db:"description"`
	Count       int64  `json:"count" db:"count"`
}

// FollowReq corresponds to the FollowReq message
type FollowReq struct {
	FollowerID  string `json:"follower_id" db:"follower_id"`
	FollowingID string `json:"following_id" db:"following_id"`
}

// FollowRes corresponds to the FollowRes message
type FollowRes struct {
	FollowerID  string `json:"follower_id" db:"follower_id"`
	FollowingID string `json:"following_id" db:"following_id"`
	FollowedAt  string `json:"followed_at" db:"followed_at"`
}
