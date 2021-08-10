package configs

import (
	"fmt"
	"pokemontrainer/models/trainers"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB to access DB
var DB *gorm.DB

type config struct {
	User      string
	Passwd    string
	Addr      string
	Port      string
	DBName    string
	Params    map[string]string
	Loc       *time.Location
	ParseTime bool
}

// InitDB initiate DB
func InitDB() {
	var db = config{
		User:      viper.Get("DB_USERNAME").(string),
		Passwd:    viper.Get("DB_PASSWORD").(string),
		Addr:      viper.Get("BASE_URL").(string),
		Port:      viper.Get("DB_PORT").(string),
		DBName:    viper.Get("DB_NAME").(string),
		Params:    map[string]string{"charset": "utf8mb4"},
		ParseTime: true,
		Loc:       time.Local,
	}

	// "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		db.User, db.Passwd, db.Addr, db.Port, db.DBName, db.Params["charset"], strconv.FormatBool(db.ParseTime), db.Loc)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
	}
	Migrate()
}

//Migrate will automigrate database from struct
func Migrate() {
	DB.AutoMigrate(&trainers.Trainer{})
}
