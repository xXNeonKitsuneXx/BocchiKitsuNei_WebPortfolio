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

type EditProjectResponse struct {
	ProjectID          *uint   `json:"project_id" validate:"required"`
	UserID             *uint   `json:"user_id" validate:"required"`
	ProjectName        *string `json:"project_name" validate:"required"`
	ProjectTag         *string `json:"project_tag" validate:"required"`
	ProjectLanguage    *string `json:"project_language" validate:"required"`
	ProjectGitHubURL   *string `json:"project_github_url" validate:"required"`
	ProjectDescription *string `json:"project_description" validate:"required"`
	ProjectPicture     *string `json:"project_picture" validate:"required"`
}

type EditProjectRequest struct {
	ProjectName        *string `json:"project_name" validate:"required"`
	ProjectTag         *string `json:"project_tag" validate:"required"`
	ProjectLanguage    *string `json:"project_language" validate:"required"`
	ProjectGitHubURL   *string `json:"project_github_url" validate:"required"`
	ProjectDescription *string `json:"project_description" validate:"required"`
	ProjectPicture     *string `json:"project_picture" validate:"required"`
}
