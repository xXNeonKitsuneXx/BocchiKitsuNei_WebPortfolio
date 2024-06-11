package dtos

type UsersResponse struct {
	UserID   *uint   `json:"user_id" validate:"required"`
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

type UserParamsResponse struct {
	UserID   *uint   `json:"user_id" validate:"required"`
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
}
