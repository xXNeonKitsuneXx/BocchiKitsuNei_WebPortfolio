package service

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/entities"
)

type ProjectService interface {
	GetProjects() ([]entities.Project, error)
	GetProjectId(int) (*entities.Project, error)
	///////////////////////////////////////////////////////////

	CreateAddProject(int, dtos.AddProjectRequest) (*entities.Project, error)
	GetProjectsFirstFour() ([]entities.Project, error)

	///////////////////////////////////////////////////////////
	//GetWishlistsOfCurrentUser(int) ([]entities.Wishlist, error)
	//GetFriendsWishlists(int) ([]entities.Wishlist, error)
	//GetWishlistDetails(int) (*entities.Wishlist, error)
	//GetProfileFriendWishlists(int, int) ([]entities.Wishlist, error)
	//
	//UpdateGrantForFriend(int, int) (*entities.Wishlist, error)
	//UpdateReceiverGotIt(int, int) (*entities.Wishlist, error)
	//UpdateReceiverDidntGetIt(int, int) (*entities.Wishlist, error)
}
