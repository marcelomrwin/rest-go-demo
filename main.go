package main

import (
	"log"
	"net/http"
	"rest-go-demo/controllers"
	"rest-go-demo/database"
	"rest-go-demo/model"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
	"gitlab.com/avarf/getenvs"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePersonByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: getenvs.GetEnvString("MYSQL_HOST", "localhost") + ":3306",
			User:       getenvs.GetEnvString("MYSQL_USER", "demo"),
			Password:   getenvs.GetEnvString("MYSQL_PASSWD", "demo"),
			DB:         getenvs.GetEnvString("MYSQL_DB", "demo"),
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&model.Person{})
}
