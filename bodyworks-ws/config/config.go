package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	driver = "mysql"
	dbUser = os.Getenv("DBUSER")
	dbPass = os.Getenv("DBPASS")
	dbHost = os.Getenv("DBHOST")
	dbPort = os.Getenv("DBPORT")
	dbName = os.Getenv("DBNAME")
)

func ConnectDB() *sql.DB {
	var stringConnexion = GetStringConnexion(dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open(driver, stringConnexion)
	if err != nil {
		log.Print(err)
		return nil
	}
	return db
}

func GetStringConnexion(user string, pass string, host string, port string, name string) string {
	return user+":"+pass+"@tcp("+host+")/"+name+"?charset=utf8&parseTime=True&loc=Local&interpolateParams=true"
}
