package main

import (
	"Stepuha.net/handler"
	"Stepuha.net/infrastructure"
	"Stepuha.net/service"
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

	db, err := infrastructure.NewPostrgesDB(infrastructure.DbConfig{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	repos := infrastructure.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	fmt.Println(pq.Elog)
	if err != nil {
		log.Fatalf("Failed to initialize DB %s", err.Error())
	}
	router := gin.Default()
	handlers.RegisterRoutes(router)
	err = router.Run(viper.GetString("port"))
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
