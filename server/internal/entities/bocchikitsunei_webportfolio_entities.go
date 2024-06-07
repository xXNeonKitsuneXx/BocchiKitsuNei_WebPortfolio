package entities

type User struct {
	UserID   *uint `gorm:"primaryKey;autoIncrement"`
	Username *string
	Password *string

	// Define associations
	Projects []Project `gorm:"foreignKey:UserID"`
}

type Project struct {
	ProjectID          *uint `gorm:"primaryKey;autoIncrement"`
	UserID             *uint `gorm:"not null"`
	ProjectName        *string
	ProjectTag         *string
	ProjectLanguage    *string
	ProjectGitHubURL   *string
	ProjectDescription *string
	ProjectPicture     *string
}
