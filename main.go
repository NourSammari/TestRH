package main

import (
	"Test/company"
	"Test/notifications"

	//"Test/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=GestionRH password=pass sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	company.InitializeDB(db)
	notifications.InitializeDB(db)
	db.AutoMigrate(&company.Company{})
	db.AutoMigrate(&notifications.Notification{})

	router := gin.Default()

	router.Use(cors.Default())

	v1 := router.Group("/api")
	{
		companies := v1.Group("/companies")
		{
			companies.GET("", company.GetCompanies)
			companies.GET("/:id", company.GetCompanyByID)
			companies.POST("", company.CreateCompany)
			companies.PUT("/:id", company.UpdateCompany)
			companies.DELETE("/:id", company.DeleteCompany)
		}

		notif := v1.Group("/notifications")
		{
			notif.GET("", notifications.GetNotifications)
			notif.GET("/:id", notifications.GetNotificationByID)
			notif.POST("", notifications.CreateNotification)
			notif.PUT("/:id", notifications.UpdateNotification)
			notif.DELETE("/:id", notifications.DeleteNotification)
		}

		/* u := v1.Group("/user")
		{
			u.GET("", users.GetUsers())
			u.GET("/:id", users.GetUser())
			u.POST("", users.Signup())
			//u.PUT("/:id", users.)
			//u.DELETE("/:id", notifications.DeleteNotification)
		} */
	}

	router.Run(":8080")
}
