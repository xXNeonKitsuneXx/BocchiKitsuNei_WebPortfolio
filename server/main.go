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
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


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

	emailService := service.NewEmailService("BocchiKitsuNei@gmail.com", "ngre hovf xnww vbkr", "BocchiKitsuNei@gmail.com", "smtp.gmail.com")
	emailHandler := handler.NewEmailHandler(emailService)

	userRepositoryDB := repository.NewUserRepositoryDB(db)
	projectRepositoryDB := repository.NewProjectRepositoryDB(db)

	userService := service.NewUserService(userRepositoryDB)
	projectService := service.NewProjectService(projectRepositoryDB)
	uploadService := service.NewUploadService(minioClient)

	userHandler := handler.NewUserHandler(userService, uploadService)
	projectHandler := handler.NewProjectHandler(projectService, uploadService)

	

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://www.bocchikitsunei.com, https://minio.bocchikitsunei.com, http://localhost:8888",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))


	//Endpoint ###########################################################################

	app.Post("/upload", storageHandler.UploadFile)

	app.Get("/GetUsers", userHandler.GetUsers)
	app.Get("/GetUserParams/:UserID", userHandler.GetUserParams)

	app.Get("/GetProjects", projectHandler.GetProjects)
	app.Get("/GetProject/:ProjectID", projectHandler.GetProjectById)
	app.Get("/GetProjectsFirstFour", projectHandler.GetProjectsFirstFour)


	app.Post("/PostAddProject/:UserID", projectHandler.AddProject) //#

	app.Get("/GetEditProject/:ProjectID", projectHandler.GetEditProject)
	app.Patch("/PatchEditProject/:ProjectID", projectHandler.UpdateEditProject)

	app.Delete("/DeleteProject/:ProjectID", projectHandler.DeleteProject)

	app.Post("/SendMail", emailHandler.SendMail)

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

