package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arfandidts/dts-be-pengenalan-microservice/menu-service/config"
	"github.com/arfandidts/dts-be-pengenalan-microservice/menu-service/database"
	"github.com/arfandidts/dts-be-pengenalan-microservice/menu-service/handler"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	con := config.Config{
		Database: config.Database{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "",
			DbName:   "dts_microservice",
			Config:   "charset=utf8&parseTime=True&loc=Local",
		},
		AuthService: config.AuthService{
			Host: "http://localhost:5001",
		},
	}

	db, err := initDB(con.Database)
	if err != nil {
		log.Panic(err)
		return
	}

	router := mux.NewRouter()

	authMiddleware := handler.AuthMiddleware{
		AuthService: con.AuthService,
	}

	menuHandler := handler.Menu{
		Db: db,
	}

	// router.Handle("/add-menu", http.HandlerFunc(menuHandler.AddMenu))
	router.Handle("/add-menu", authMiddleware.ValidateAuth(http.HandlerFunc(menuHandler.AddMenu)))
	router.Handle("/menu", http.HandlerFunc(menuHandler.GetAllMenu))

	fmt.Println("Menu service listen on port :5000")
	log.Panic(http.ListenAndServe(":5000", router))
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Membuat tabel jika belum ada
	err = db.AutoMigrate(database.Menu{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to DB")

	return db, nil
}
