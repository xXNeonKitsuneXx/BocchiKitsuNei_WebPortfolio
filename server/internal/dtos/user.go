package dtos

type UsersResponse struct {
	UserID   *uint   `json:"user_id" validate:"required"`
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

type UserIDResponse struct {
	UserID   *uint   `json:"user_id" validate:"required"`
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

type UserParamsResponse struct {
	UserID   *uint   `json:"user_id" validate:"required"`
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

//////////////////////////////////////////////////////////////////////////////

type EditUserProfileResponse struct {
	UserID    *uint   `json:"user_id" validate:"required"`
	Username  *string `json:"username" validate:"required"`
	Email     *string `json:"email" validate:"required"`
	Firstname *string `json:"firstname" validate:"required"`
	Lastname  *string `json:"lastname" validate:"required"`
	PhoneNum  *string `json:"phone_num" validate:"required"`
}

type EditUserProfileRequest struct {
	Username  *string `json:"username" validate:"required"`
	Email     *string `json:"email" validate:"required"`
	Firstname *string `json:"firstname" validate:"required"`
	Lastname  *string `json:"lastname" validate:"required"`
	PhoneNum  *string `json:"phone_num" validate:"required"`
}

type RegisterRequest struct {
	Username  *string `json:"username" validate:"required"`
	Password  *string `json:"password" validate:"required"`
	Email     *string `json:"email" validate:"required"`
	Firstname *string `json:"firstname" validate:"required"`
	Lastname  *string `json:"lastname" validate:"required"`
	PhoneNum  *string `json:"phone_num" validate:"required"`
	UserPic   *string `json:"user_pic" form:"userPic" validate:"required"`
}

type LoginRequest struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

type UserInfoResponse struct {
	UserID   *uint   `json:"user_id" validate:"required"`
	Username *string `json:"username" validate:"required"`
	UserPic  *string `json:"user_pic" validate:"required"`
	Token    *string `json:"token,omitempty"`
}

type LoginResponse struct {
	UserID   *uint   `json:"user_id" validate:"required"`
	Username *string `json:"username" validate:"required"`
	Token    *string `json:"token,omitempty"`
}
