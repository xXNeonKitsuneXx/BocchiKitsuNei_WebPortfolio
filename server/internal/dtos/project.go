package dtos

type ProjectsResponse struct {
	ProjectID          *uint   `json:"project_id" validate:"required"`
	UserID             *uint   `json:"user_id" validate:"required"`
	ProjectName        *string `json:"project_name" validate:"required"`
	ProjectTag         *string `json:"project_tag" validate:"required"`
	ProjectLanguage    *string `json:"project_language" validate:"required"`
	ProjectGitHubURL   *string `json:"project_github_url" validate:"required"`
	ProjectDescription *string `json:"project_description" validate:"required"`
	ProjectPicture     *string `json:"project_picture" validate:"required"`
}

type ProjectParamsResponse struct {
	ProjectID          *uint   `json:"project_id" validate:"required"`
	UserID             *uint   `json:"user_id" validate:"required"`
	ProjectName        *string `json:"project_name" validate:"required"`
	ProjectTag         *string `json:"project_tag" validate:"required"`
	ProjectLanguage    *string `json:"project_language" validate:"required"`
	ProjectGitHubURL   *string `json:"project_github_url" validate:"required"`
	ProjectDescription *string `json:"project_description" validate:"required"`
	ProjectPicture     *string `json:"project_picture" validate:"required"`
}

type AddProjectRequest struct {
	UserID             *uint   `json:"user_id" validate:"required"`
	ProjectName        *string `json:"project_name" validate:"required"`
	ProjectTag         *string `json:"project_tag" validate:"required"`
	ProjectLanguage    *string `json:"project_language" validate:"required"`
	ProjectGitHubURL   *string `json:"project_github_url" validate:"required"`
	ProjectDescription *string `json:"project_description" validate:"required"`
	ProjectPicture     *string `json:"project_picture" validate:"required"`
}

type ProjectsFirstFourResponse struct {
	ProjectID          *uint   `json:"project_id" validate:"required"`
	UserID             *uint   `json:"user_id" validate:"required"`
	ProjectName        *string `json:"project_name" validate:"required"`
	ProjectTag         *string `json:"project_tag" validate:"required"`
	ProjectLanguage    *string `json:"project_language" validate:"required"`
	ProjectGitHubURL   *string `json:"project_github_url" validate:"required"`
	ProjectDescription *string `json:"project_description" validate:"required"`
	ProjectPicture     *string `json:"project_picture" validate:"required"`
}

//type WishlistsUserCurrentResponse struct {
//	WishlistID      *uint   `json:"wishlist_id" validate:"required"`
//	UserID          *uint   `json:"user_id" validate:"required"`
//	Itemname        *string `json:"itemname" validate:"required"`
//	Price           *uint   `json:"price" validate:"required"`
//	LinkURL         *string `json:"link_url" validate:"required"`
//	ItemPic         *string `json:"item_pic" validate:"required"`
//	AlreadyBought   *bool   `json:"already_bought" validate:"required"`
//	GrantedByUserID *uint   `json:"granted_by_user_id" validate:"required"`
//}
//
//type WishlistIDInfoResponse struct {
//	WishlistID      *uint   `json:"wishlist_id" validate:"required"`
//	UserID          *uint   `json:"user_id" validate:"required"`
//	Itemname        *string `json:"itemname" validate:"required"`
//	Price           *uint   `json:"price" validate:"required"`
//	LinkURL         *string `json:"link_url" validate:"required"`
//	ItemPic         *string `json:"item_pic" validate:"required"`
//	AlreadyBought   *bool   `json:"already_bought" validate:"required"`
//	GrantedByUserID *uint   `json:"granted_by_user_id" validate:"required"`
//}
//
/////////////////////////////////////////////////////
//
//type WishlistsOfCurrentUserResponse struct {
//	WishlistID        *uint   `json:"wishlist_id" validate:"required"`
//	UserID            *uint   `json:"user_id" validate:"required"`
//	Itemname          *string `json:"itemname" validate:"required"`
//	Price             *uint   `json:"price" validate:"required"`
//	LinkURL           *string `json:"link_url" validate:"required"`
//	ItemPic           *string `json:"item_pic" validate:"required"`
//	AlreadyBought     *bool   `json:"already_bought" validate:"required"`
//	GrantedByUserID   *uint   `json:"granted_by_user_id" validate:"required"`
//	UsernameOfGranter *string `json:"username_of_granter" validate:"required"`
//}
//
//type FriendsWishlistsResponse struct {
//	WishlistID         *uint   `json:"wishlist_id" validate:"required"`
//	UserID             *uint   `json:"user_id" validate:"required"`
//	Itemname           *string `json:"itemname" validate:"required"`
//	Price              *uint   `json:"price" validate:"required"`
//	LinkURL            *string `json:"link_url" validate:"required"`
//	ItemPic            *string `json:"item_pic" validate:"required"`
//	AlreadyBought      *bool   `json:"already_bought" validate:"required"`
//	GrantedByUserID    *uint   `json:"granted_by_user_id" validate:"required"`
//	UsernameOfWishlist *string `json:"username_of_wishlist" validate:"required"`
//	UserPicOfWishlist  *string `json:"user_pic_of_wishlist" validate:"required"`
//}
//
//type WishlistDetailsResponse struct {
//	WishlistID      *uint   `json:"wishlist_id" validate:"required"`
//	UserID          *uint   `json:"user_id" validate:"required"`
//	Itemname        *string `json:"itemname" validate:"required"`
//	Price           *uint   `json:"price" validate:"required"`
//	LinkURL         *string `json:"link_url" validate:"required"`
//	ItemPic         *string `json:"item_pic" validate:"required"`
//	AlreadyBought   *bool   `json:"already_bought" validate:"required"`
//	GrantedByUserID *uint   `json:"granted_by_user_id" validate:"required"`
//}
//
//type ProfileFriendWishlistsResponse struct {
//	WishlistID         *uint   `json:"wishlist_id" validate:"required"`
//	UserID             *uint   `json:"user_id" validate:"required"`
//	Itemname           *string `json:"itemname" validate:"required"`
//	Price              *uint   `json:"price" validate:"required"`
//	LinkURL            *string `json:"link_url" validate:"required"`
//	ItemPic            *string `json:"item_pic" validate:"required"`
//	AlreadyBought      *bool   `json:"already_bought" validate:"required"`
//	GrantedByUserID    *uint   `json:"granted_by_user_id" validate:"required"`
//	UsernameOfWishlist *string `json:"username_of_wishlist" validate:"required"`
//	UserPicOfWishlist  *string `json:"user_pic_of_wishlist" validate:"required"`
//}
