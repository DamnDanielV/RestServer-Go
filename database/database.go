package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

// connectDB crea una conexion con la base de datos
// si se genera un error muestra el mensaje en la consola
func connectDB() *gorm.DB {
	var err error
	const DB_USERNAME = "root"
	MYSQL_ROOT_PASSWORD := os.Getenv("MYSQL_ROOT_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	dsn := DB_USERNAME + ":" + MYSQL_ROOT_PASSWORD + "@tcp" + "(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
