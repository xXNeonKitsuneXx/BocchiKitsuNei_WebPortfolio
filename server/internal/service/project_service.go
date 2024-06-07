package service

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/entities"
	"bocchikitsunei_webportfolio/internal/repository"
	"bocchikitsunei_webportfolio/internal/utils/v"
	"log"
)

type projectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) projectService {
	return projectService{projectRepo: projectRepo}
}

func (s projectService) GetProjects() ([]entities.Project, error) {
	projects, err := s.projectRepo.GetAllProject()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	projectResponses := []entities.Project{}
	for _, project := range projects {
		projectResponse := entities.Project{
			ProjectID:          project.ProjectID,
			UserID:             project.UserID,
			ProjectName:        project.ProjectName,
			ProjectTag:         project.ProjectTag,
			ProjectLanguage:    project.ProjectLanguage,
			ProjectGitHubURL:   project.ProjectGitHubURL,
			ProjectDescription: project.ProjectDescription,
			ProjectPicture:     project.ProjectPicture,
		}
		projectResponses = append(projectResponses, projectResponse)
	}
	return projectResponses, nil
}

func (s projectService) GetProjectId(projectid int) (*entities.Project, error) {
	project, err := s.projectRepo.GetProjectById(projectid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	projectResponse := entities.Project{
		ProjectID:          project.ProjectID,
		UserID:             project.UserID,
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}
	return &projectResponse, nil
}

func (s projectService) CreateAddProject(userid int, req dtos.AddProjectRequest) (*entities.Project, error) {
	project := &entities.Project{
		UserID:             v.UintPtr(userid),
		ProjectName:        req.ProjectName,
		ProjectTag:         req.ProjectTag,
		ProjectLanguage:    req.ProjectLanguage,
		ProjectGitHubURL:   req.ProjectGitHubURL,
		ProjectDescription: req.ProjectDescription,
		ProjectPicture:     req.ProjectPicture,
	}

	err := s.projectRepo.PostAddProject(project)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return project, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

//func (s wishlistService) GetWishlistsOfCurrentUser(userid int) ([]entities.Wishlist, error) {
//	wishlists, err := s.wishlistRepo.GetAllWishlistsOfCurrentUserId(userid)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	wishlistResponses := []entities.Wishlist{}
//	for _, wishlist := range wishlists {
//		wishlistResponse := entities.Wishlist{
//			WishlistID:        wishlist.WishlistID,
//			UserID:            wishlist.UserID,
//			Itemname:          wishlist.Itemname,
//			Price:             wishlist.Price,
//			LinkURL:           wishlist.LinkURL,
//			ItemPic:           wishlist.ItemPic,
//			AlreadyBought:     wishlist.AlreadyBought,
//			GrantedByUserID:   wishlist.GrantedByUserID,
//			UsernameOfGranter: wishlist.UsernameOfGranter,
//		}
//		wishlistResponses = append(wishlistResponses, wishlistResponse)
//	}
//	return wishlistResponses, nil
//}
//
//func (s wishlistService) GetFriendsWishlists(userid int) ([]entities.Wishlist, error) {
//	wishlists, err := s.wishlistRepo.GetAllFriendsWishlists(userid)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	wishlistResponses := []entities.Wishlist{}
//	for _, wishlist := range wishlists {
//		wishlistResponse := entities.Wishlist{
//			WishlistID:         wishlist.WishlistID,
//			UserID:             wishlist.UserID,
//			Itemname:           wishlist.Itemname,
//			Price:              wishlist.Price,
//			LinkURL:            wishlist.LinkURL,
//			ItemPic:            wishlist.ItemPic,
//			AlreadyBought:      wishlist.AlreadyBought,
//			GrantedByUserID:    wishlist.GrantedByUserID,
//			UsernameOfWishlist: wishlist.UsernameOfWishlist,
//			UserPicOfWishlist:  wishlist.UserPicOfWishlist,
//		}
//		wishlistResponses = append(wishlistResponses, wishlistResponse)
//	}
//	return wishlistResponses, nil
//}
//
//func (s wishlistService) GetWishlistDetails(wishlistid int) (*entities.Wishlist, error) {
//	wishlist, err := s.wishlistRepo.GetWishlistDetailsByWishlistId(wishlistid)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	wishlistResponse := entities.Wishlist{
//		WishlistID:      wishlist.WishlistID,
//		UserID:          wishlist.UserID,
//		Itemname:        wishlist.Itemname,
//		Price:           wishlist.Price,
//		LinkURL:         wishlist.LinkURL,
//		ItemPic:         wishlist.ItemPic,
//		AlreadyBought:   wishlist.AlreadyBought,
//		GrantedByUserID: wishlist.GrantedByUserID,
//	}
//	return &wishlistResponse, nil
//}
//
//func (s wishlistService) GetProfileFriendWishlists(currentUserID, wishlistOwnerID int) ([]entities.Wishlist, error) {
//	wishlists, err := s.wishlistRepo.GetAllProfileFriendWishlists(currentUserID, wishlistOwnerID)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	wishlistResponses := []entities.Wishlist{}
//	for _, wishlist := range wishlists {
//		wishlistResponse := entities.Wishlist{
//			WishlistID:         wishlist.WishlistID,
//			UserID:             wishlist.UserID,
//			Itemname:           wishlist.Itemname,
//			Price:              wishlist.Price,
//			LinkURL:            wishlist.LinkURL,
//			ItemPic:            wishlist.ItemPic,
//			AlreadyBought:      wishlist.AlreadyBought,
//			GrantedByUserID:    wishlist.GrantedByUserID,
//			UsernameOfWishlist: wishlist.UsernameOfWishlist,
//			UserPicOfWishlist:  wishlist.UserPicOfWishlist,
//		}
//		wishlistResponses = append(wishlistResponses, wishlistResponse)
//	}
//	return wishlistResponses, nil
//}
//
//func (s wishlistService) UpdateGrantForFriend(wishlistID, granterUserID int) (*entities.Wishlist, error) {
//	wishlist, err := s.wishlistRepo.GetWishlistByWishlistId(wishlistID)
//	if err != nil {
//		return nil, err
//	}
//
//	bought := false
//	wishlist.AlreadyBought = &bought
//
//	wishlist.GrantedByUserID = v.UintPtr(granterUserID)
//
//	err = s.wishlistRepo.UpdateGrantForFriend(wishlist)
//	if err != nil {
//		return nil, err
//	}
//
//	return wishlist, nil
//}
//
//func (s wishlistService) UpdateReceiverGotIt(wishlistID, granterUserID int) (*entities.Wishlist, error) {
//	wishlist, err := s.wishlistRepo.GetWishlistByWishlistId(wishlistID)
//	if err != nil {
//		return nil, err
//	}
//
//	bought := true
//	wishlist.AlreadyBought = &bought
//
//	wishlist.GrantedByUserID = v.UintPtr(granterUserID)
//
//	err = s.wishlistRepo.UpdateReceiverGotIt(wishlist)
//	if err != nil {
//		return nil, err
//	}
//
//	return wishlist, nil
//}
//
//func (s wishlistService) UpdateReceiverDidntGetIt(wishlistID, granterUserID int) (*entities.Wishlist, error) {
//	wishlist, err := s.wishlistRepo.GetWishlistByWishlistId(wishlistID)
//	if err != nil {
//		return nil, err
//	}
//
//	var bought *bool = nil
//	wishlist.AlreadyBought = bought
//
//	wishlist.GrantedByUserID = v.UintPtr(granterUserID)
//
//	err = s.wishlistRepo.UpdateReceiverDidntGetIt(wishlist)
//	if err != nil {
//		return nil, err
//	}
//
//	return wishlist, nil
//}
//
