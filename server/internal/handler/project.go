package handler

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type projectHandler struct {
	projectSer service.ProjectService
	jwtSecret  string
	uploadSer  service.UploadService
}

func NewProjectHandler(projectSer service.ProjectService, jwtSecret string, uploadSer service.UploadService) projectHandler {
	return projectHandler{projectSer: projectSer, jwtSecret: jwtSecret, uploadSer: uploadSer}
}

func (h *projectHandler) GetProjects(c *fiber.Ctx) error {
	projectsResponse := make([]dtos.ProjectsResponse, 0)

	projects, err := h.projectSer.GetProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		projectsResponse = append(projectsResponse, dtos.ProjectsResponse{
			ProjectID:          project.ProjectID,
			UserID:             project.UserID,
			ProjectName:        project.ProjectName,
			ProjectTag:         project.ProjectTag,
			ProjectLanguage:    project.ProjectLanguage,
			ProjectGitHubURL:   project.ProjectGitHubURL,
			ProjectDescription: project.ProjectDescription,
			ProjectPicture:     project.ProjectPicture,
		})
	}
	return c.JSON(projectsResponse)
}

func (h *projectHandler) GetProjectById(c *fiber.Ctx) error {

	projectIDReceive, err := strconv.Atoi(c.Params("ProjectID"))

	project, err := h.projectSer.GetProjectId(projectIDReceive)
	if err != nil {
		return err
	}

	projectResponse := dtos.ProjectsResponse{
		ProjectID:          project.ProjectID,
		UserID:             project.UserID,
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}

	return c.JSON(projectResponse)
}

func (h *projectHandler) AddProject(c *fiber.Ctx) error {
	//// Extract the token from the request headers
	//token := c.Get("Authorization")
	//
	//// Check if the token is empty
	//if token == "" {
	//	return errors.New("token is missing")
	//}
	//
	//// Extract the user ID from the token
	//userIDExtract, err := utils.ExtractUserIDFromToken(strings.Replace(token, "Bearer ", "", 1), h.jwtSecret)
	//if err != nil {
	//	return err
	//}

	userIDExtract, err := strconv.Atoi(c.Params("UserID"))

	var request dtos.AddProjectRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	// Check if a file is uploaded
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "File not found")
	}

	// Call upload service to upload the file
	fileURL, err := h.uploadSer.UploadFile(file)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload file")
	}

	// Set the uploaded file URL in the registration request
	request.ProjectPicture = fileURL

	// Check if project_picture field is empty or nil
	if request.ProjectPicture == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Project picture is required")
	}

	project, err := h.projectSer.CreateAddProject(userIDExtract, request)
	if err != nil {
		return err
	}

	projectResponse := dtos.AddProjectRequest{
		UserID:             project.UserID,
		ProjectName:        project.ProjectName,
		ProjectTag:         project.ProjectTag,
		ProjectLanguage:    project.ProjectLanguage,
		ProjectGitHubURL:   project.ProjectGitHubURL,
		ProjectDescription: project.ProjectDescription,
		ProjectPicture:     project.ProjectPicture,
	}

	return c.JSON(projectResponse)
}

func (h *projectHandler) GetProjectsFirstFour(c *fiber.Ctx) error {
	projectsResponse := make([]dtos.ProjectsFirstFourResponse, 0)

	projects, err := h.projectSer.GetProjectsFirstFour()
	if err != nil {
		return err
	}

	for _, project := range projects {
		projectsResponse = append(projectsResponse, dtos.ProjectsFirstFourResponse{
			ProjectID:          project.ProjectID,
			UserID:             project.UserID,
			ProjectName:        project.ProjectName,
			ProjectTag:         project.ProjectTag,
			ProjectLanguage:    project.ProjectLanguage,
			ProjectGitHubURL:   project.ProjectGitHubURL,
			ProjectDescription: project.ProjectDescription,
			ProjectPicture:     project.ProjectPicture,
		})
	}
	return c.JSON(projectsResponse)
}

//****************************************************************************

//func (h *wishlistHandler) GetWishlistsOfCurrentUser(c *fiber.Ctx) error {
//	// Extract the token from the request headers
//	token := c.Get("Authorization")
//
//	// Check if the token is empty
//	if token == "" {
//		return errors.New("token is missing")
//	}
//
//	// Extract the user ID from the token
//	userID, err := utils.ExtractUserIDFromToken(strings.Replace(token, "Bearer ", "", 1), h.jwtSecret)
//	if err != nil {
//		return err
//	}
//
//	wishlists, err := h.wishlistSer.GetWishlistsOfCurrentUser(userID)
//	if err != nil {
//		return err
//	}
//
//	wishlistsResponse := make([]dtos.WishlistsOfCurrentUserResponse, 0)
//	for _, wishlist := range wishlists {
//		wishlistsResponse = append(wishlistsResponse, dtos.WishlistsOfCurrentUserResponse{
//			WishlistID:        wishlist.WishlistID,
//			UserID:            wishlist.UserID,
//			Itemname:          wishlist.Itemname,
//			Price:             wishlist.Price,
//			LinkURL:           wishlist.LinkURL,
//			ItemPic:           wishlist.ItemPic,
//			AlreadyBought:     wishlist.AlreadyBought,
//			GrantedByUserID:   wishlist.GrantedByUserID,
//			UsernameOfGranter: wishlist.UsernameOfGranter,
//		})
//	}
//	return c.JSON(wishlistsResponse)
//}
//
//func (h *wishlistHandler) GetFriendsWishlists(c *fiber.Ctx) error {
//
//	// Extract the token from the request headers
//	token := c.Get("Authorization")
//
//	// Check if the token is empty
//	if token == "" {
//		return errors.New("token is missing")
//	}
//
//	// Extract the user ID from the token
//	userIDExtract, err := utils.ExtractUserIDFromToken(strings.Replace(token, "Bearer ", "", 1), h.jwtSecret)
//	if err != nil {
//		return err
//	}
//
//	wishlistsResponse := make([]dtos.FriendsWishlistsResponse, 0)
//	wishlists, err := h.wishlistSer.GetFriendsWishlists(userIDExtract)
//	if err != nil {
//		return err
//	}
//
//	for _, wishlist := range wishlists {
//		wishlistsResponse = append(wishlistsResponse, dtos.FriendsWishlistsResponse{
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
//		})
//	}
//	return c.JSON(wishlistsResponse)
//}
//
//func (h *wishlistHandler) GetWishlistDetails(c *fiber.Ctx) error {
//	wishlistID := c.Params("wishlistID")
//
//	wishlistIDReceive, err := strconv.Atoi(wishlistID)
//	if err != nil {
//		return err
//	}
//
//	wishlist, err := h.wishlistSer.GetWishlistDetails(wishlistIDReceive)
//	if err != nil {
//		return err
//	}
//
//	wishlistResponse := dtos.WishlistDetailsResponse{
//		WishlistID:      wishlist.WishlistID,
//		UserID:          wishlist.UserID,
//		Itemname:        wishlist.Itemname,
//		Price:           wishlist.Price,
//		LinkURL:         wishlist.LinkURL,
//		ItemPic:         wishlist.ItemPic,
//		AlreadyBought:   wishlist.AlreadyBought,
//		GrantedByUserID: wishlist.GrantedByUserID,
//	}
//
//	return c.JSON(wishlistResponse)
//}
//
//func (h *wishlistHandler) GetProfileFriendWishlists(c *fiber.Ctx) error {
//
//	currentUserID, err := strconv.Atoi(c.Params("CurrentUserID"))
//	if err != nil {
//		return err
//	}
//
//	wishlistOwnerID, err := strconv.Atoi(c.Params("WishlistOwnerID"))
//	if err != nil {
//		return err
//	}
//
//	wishlistsResponse := make([]dtos.FriendsWishlistsResponse, 0)
//	wishlists, err := h.wishlistSer.GetProfileFriendWishlists(currentUserID, wishlistOwnerID)
//	if err != nil {
//		return err
//	}
//
//	for _, wishlist := range wishlists {
//		wishlistsResponse = append(wishlistsResponse, dtos.FriendsWishlistsResponse{
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
//		})
//	}
//	return c.JSON(wishlistsResponse)
//}
//
//func (h *wishlistHandler) UpdateGrantForFriend(c *fiber.Ctx) error {
//	wishlistID, err := strconv.Atoi(c.Params("WishlistID"))
//	if err != nil {
//		return err
//	}
//
//	granterUserID, err := strconv.Atoi(c.Params("GranterUserID"))
//	if err != nil {
//		return err
//	}
//
//	_, err = h.wishlistSer.UpdateGrantForFriend(wishlistID, granterUserID)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(fiber.Map{"message": "UpdateGrantForFriend successfully"})
//}
//
//func (h *wishlistHandler) UpdateReceiverGotIt(c *fiber.Ctx) error {
//	wishlistID, err := strconv.Atoi(c.Params("WishlistID"))
//	if err != nil {
//		return err
//	}
//
//	granterUserID, err := strconv.Atoi(c.Params("GranterUserID"))
//	if err != nil {
//		return err
//	}
//
//	_, err = h.wishlistSer.UpdateReceiverGotIt(wishlistID, granterUserID)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(fiber.Map{"message": "UpdateReceiverGotIt successfully"})
//}
//
//func (h *wishlistHandler) UpdateReceiverDidntGetIt(c *fiber.Ctx) error {
//	wishlistID, err := strconv.Atoi(c.Params("WishlistID"))
//	if err != nil {
//		return err
//	}
//
//	granterUserID, err := strconv.Atoi(c.Params("GranterUserID"))
//	if err != nil {
//		return err
//	}
//
//	_, err = h.wishlistSer.UpdateReceiverDidntGetIt(wishlistID, granterUserID)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(fiber.Map{"message": "UpdateReceiverDidntGetIt successfully"})
//}
//
