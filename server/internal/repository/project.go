package repository

import "bocchikitsunei_webportfolio/internal/entities"

type ProjectRepository interface {
	GetAllProject() ([]entities.Project, error)
	GetProjectById(int) (*entities.Project, error)

	//////////////////////////////////////////////////////////////

	PostAddProject(project *entities.Project) error
	GetProjectsFirstFour() ([]entities.Project, error)

	////////////////////////////////////////////////////////////////////

	//GetAllWishlistsOfCurrentUserId(int) ([]entities.Project, error)
	//GetAllFriendsWishlists(int) ([]entities.Project, error)
	//GetWishlistDetailsByWishlistId(int) (*entities.Project, error)
	//GetAllProfileFriendWishlists(int, int) ([]entities.Project, error)
	//
	//UpdateGrantForFriend(wishlist *entities.Project) error
	//UpdateReceiverGotIt(wishlist *entities.Project) error
	//UpdateReceiverDidntGetIt(wishlist *entities.Project) error

}
