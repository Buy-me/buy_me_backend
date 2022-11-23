package main

import (
	"food_delivery/component/appctx"
	"food_delivery/component/uploadprovider"
	"food_delivery/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"` // tag
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {

	err := godotenv.Load()

	dsn := os.Getenv("MYSQL_CONNECTION")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	log.Println(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provider, secretKey)

	// **************** DEMO GIN REACT API *********************
	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("/static", "./static")
	//****** POST ******
	v1 := r.Group("/v1")

	setUpRoutes(appContext, v1)
	setUpAdminRoutes(appContext, v1)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// +++++++++++++++++++++++ DEMO GORM *******************

	// newRestaurant := Restaurant{Name: "Beff Bitch", Addr: "99 Thu Duc"}

	// if err := db.Create(&newRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println("New id: ", newRestaurant.Id)

	// var myRestaurant Restaurant

	// if err := db.Where("id = ?", 1).First(&myRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println(myRestaurant)

	// newName := "200 Lab"
	// updateData := RestaurantUpdate{Name: &newName}

	// myRestaurant.Name = ""

	// if err := db.Where("id = ?", 2).Updates(&updateData).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println(myRestaurant)

	// if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println("Delete Successfully")

}
