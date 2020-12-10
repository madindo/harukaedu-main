package bootstrap
import (
	"fmt"
	"harukaedu-main/database"
	"harukaedu-main/router"
	"github.com/jinzhu/gorm"
	"github.com/madindo/harukaedu-users/models"
)

var db *gorm.DB
var err error

func BootApplication() {
	fmt.Println("Bootstrapping the Application")
	router.Web()
	db := database.Connect()
	defer db.Close()
	db.Debug().DropTableIfExists(&models.User{})
	db.Debug().DropTableIfExists(&models.Address{})
	db.Debug().AutoMigrate(&models.User{}, &models.Address{})
}
