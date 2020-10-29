package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arfandidts/dts-be-pendalaman-microservice/auth-service/config"
	"github.com/arfandidts/dts-be-pendalaman-microservice/auth-service/handler"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	if cfg, err := getConfig(); err != nil {
		log.Println(err.Error())
	} else {
		log.Println(cfg)
	}

	router := mux.NewRouter()

	router.Handle("/admin-auth", http.HandlerFunc(handler.ValidateAuth))

	fmt.Printf("Auth service listen on :5001")
	log.Panic(http.ListenAndServe(":5001", router))
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
