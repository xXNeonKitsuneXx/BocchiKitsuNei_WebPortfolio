package main

import (
	"bocchikitsunei_webportfolio/internal/entities"
	"bocchikitsunei_webportfolio/internal/handler"
	"bocchikitsunei_webportfolio/internal/repository"
	"bocchikitsunei_webportfolio/internal/service"
	"fmt"
	"log"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const jwtSecret = "BocchiKitsuNeiSecret"

func main() {
	initTimeZone()
	initConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	log.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		panic("Failed to AutoMigrate User")
	}

	err = db.AutoMigrate(&entities.Project{})
	if err != nil {
		panic("Failed to AutoMigrate Project")
	}

	minioClient, err := minio.New(viper.GetString("minio.host")+":"+viper.GetString("minio.port"), &minio.Options{
		Creds:  credentials.NewStaticV4("EneudJTZpjRIZuFYq6MF", "eo1uLxOg33oeuOHbI7kokmSNSp1AgJTfw1QMWLda", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Minio connected")

	uploadSer := service.NewUploadService(minioClient)
	storageHandler := handler.NewStorageHandler(uploadSer)

	userRepositoryDB := repository.NewUserRepositoryDB(db)
	projectRepositoryDB := repository.NewProjectRepositoryDB(db)

	userService := service.NewUserService(userRepositoryDB, jwtSecret)
	projectService := service.NewProjectService(projectRepositoryDB)
	uploadService := service.NewUploadService(minioClient)

	userHandler := handler.NewUserHandler(userService, jwtSecret, uploadService)
	projectHandler := handler.NewProjectHandler(projectService, jwtSecret, uploadService)

	app := fiber.New()

	//app.Use(func(c *fiber.Ctx) error {
	//	if c.Path() != "/Register" && c.Path() != "/Login" {
	//		jwtMiddleware := jwtware.New(jwtware.Config{
	//			SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	//			ErrorHandler: func(c *fiber.Ctx, err error) error {
	//				return fiber.ErrUnauthorized
	//			},
	//		})
	//		return jwtMiddleware(c)
	//	}
	//	return c.Next()
	//})

	//Endpoint ###########################################################################

	// Just endpoint of test
	app.Get("/Users", userHandler.GetUsers)
	app.Get("/User", userHandler.GetUserId) //#
	app.Get("/User/:UserID", userHandler.GetUserParams)

	app.Get("/Projects", projectHandler.GetProjects)
	app.Get("/Project/:ProjectID", projectHandler.GetProjectById)

	app.Post("/upload", storageHandler.UploadFile)

	app.Post("/AddProject/:UserID", projectHandler.AddProject) //#

	//////////////////////////////////////////////////////////////////////////////////////

	// Project use these endpoints
	app.Post("/Register", userHandler.Register)
	app.Post("/Login", userHandler.Login)

	app.Get("/GetEditUserProfile/:UserID", userHandler.GetEditUserProfile)
	app.Patch("/UpdateEditUserProfile/:UserID", userHandler.UpdateEditUserProfile)

	////////////

	//app.Get("/GetWishlistsOfCurrentUser", wishlistHandler.GetWishlistsOfCurrentUser) //#
	//
	//app.Get("/GetFriendsWishlists", wishlistHandler.GetFriendsWishlists) //#
	//app.Get("/GetWishlistDetails/:WishlistID", wishlistHandler.GetWishlistDetails)
	//app.Get("/GetProfileFriendWishlists/:CurrentUserID/:WishlistOwnerID", wishlistHandler.GetProfileFriendWishlists)
	//
	//app.Put("/UpdateGrantForFriend/:WishlistID/:GranterUserID", wishlistHandler.UpdateGrantForFriend)
	//app.Put("/UpdateReceiverGotIt/:WishlistID/:GranterUserID", wishlistHandler.UpdateReceiverGotIt)
	//app.Put("/UpdateReceiverDidntGetIt/:WishlistID/:GranterUserID", wishlistHandler.UpdateReceiverDidntGetIt)

	//#####################################################################################

	log.Printf("BocchiKitsuNei WebPortfolio run at port:  %v", viper.GetInt("app.port"))
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}