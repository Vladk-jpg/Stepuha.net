package main

import (
	"Server/postgre"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Couldn't init config %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Couldn't read the enviroment variables %s", err.Error())
	}
	db, err := postgre.NewPostrgesDB(postgre.DbConfig{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	fmt.Println(pq.Elog)
	if err != nil {
		log.Fatalf("Failed to initialize DB %s", err.Error())
	}
	err = db.Close()
	r := gin.Default()
	err = r.Run(viper.GetString("port"))
	if err != nil {
		fmt.Println("Ooops...")
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
