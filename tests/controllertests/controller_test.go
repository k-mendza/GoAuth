package controllertests

import (
	"GoAuth/api/controllers"
	"GoAuth/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var server = controllers.Server{}
var userInstance = models.User{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}

func Database() {
	var err error
	TestDbDriver := os.Getenv("TEST_DB_DRIVER")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
	server.DB, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		log.Printf("cannot connect to %s database\n", TestDbDriver)
		log.Fatal("database error:", err)
	} else {
		log.Printf("sucessfuly connected to %s database", TestDbDriver)
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		FirstName: "John",
		LastName:  "Kowalski",
		Email:     "johnkowalski@gmail.com",
		Password:  "password",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() ([]models.User, error) {
	users := []models.User{
		{
			FirstName: "Steven",
			LastName:  "Victor",
			Email:     "steven@gmail.com",
			Password:  "password",
		}, {
			FirstName: "Kenny",
			LastName:  "Morris",
			Email:     "kenny@gmail.com",
			Password:  "password",
		},
	}
	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}
