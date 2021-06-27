package main

import (
	"fmt"
	mongodb "github.com/elissonalvesilva/interview-meli/api/infra/db/mongodb/helper"
	"github.com/elissonalvesilva/interview-meli/api/server/factories/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run(application string) {
	fmt.Println(application+" Is running in port", a.httpServer.Addr)
	err := a.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func NewApp(port int) *App {
	dbClient, err := mongodb.NewClient()
	if err != nil {
		panic(fmt.Errorf("Failed to create new database client: %s", err))
	}

	err = dbClient.Connect()
	if err != nil {
		panic(fmt.Errorf("Failed to connect to database: %s", err.Error()))
	}

	db := mongodb.NewDatabase(dbClient)

	router := mux.NewRouter()


	controller := controllers.MakeDNAController(db)

	router.HandleFunc("/simian", controller.AnalyzeDNA).Methods("POST")
	router.HandleFunc("/stats", controller.GetStatsDNAs).Methods("GET")

	return &App{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: router,
		},
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error to load .env file err: %v", err)
	}
}

func main() {
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	name := os.Getenv("APP_NAME")

	NewApp(port).Run(name)
}
