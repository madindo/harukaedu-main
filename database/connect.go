package database
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/madindo/harukaedu-main/errors"
)


var db *gorm.DB
var err error

func Connect() *gorm.DB {
	db, err = gorm.Open("sqlite3", "database/harukaedu.db")
	errors.ErrorConnection(err)
	return db
}