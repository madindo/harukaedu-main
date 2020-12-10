package bootstrap
import (
	"fmt"
	"harukaedu-main/router"
	"github.com/jinzhu/gorm"
	"harukaedu-main/database"
	"harukaedu-users/models"
)

var db *gorm.DB
var err error

func BootApplication() {
	fmt.Println("Bootstrapping the Application")
	router.Web()
	db := database.Connect()
	defer db.Close()
	db.Debug().AutoMigrate(&models.User{})
}
