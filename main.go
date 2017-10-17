package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/betacraft/yaag/middleware"
	"github.com/betacraft/yaag/yaag"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var appPort string
var version string

func main() {

	// Get and open a DB connections
	// EXAMPLE MYSQL_CONNECTION: username:password@tcp(dbserver.com(or IP address):3306)/databaseName
	InitDB(os.Getenv("MYSQL_CONNECTION"))
	fmt.Println("Database connection initialized")

	// This is for generating API documentation. I am not sure how I feel about it
	yaag.Init(&yaag.Config{On: false, DocTitle: "s_k-api", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "https://steve-and-kyle.appspot.com", "Staging": "iCantAffordThat.sorry"}})

	r := mux.NewRouter()
	r.HandleFunc("/api/episodes", middleware.HandleFunc(AllEps))
	r.HandleFunc("/api/episode", middleware.HandleFunc(AllEps))
	r.HandleFunc("/api/episode/{id}", middleware.HandleFunc(SingleEpisode))
	r.HandleFunc("/api/stats", middleware.HandleFunc(StatsIndex))
	r.HandleFunc("/api", middleware.HandleFunc(Index))
	fmt.Println("Routes set up")
	// Catches anything left
	http.Handle("/", r)

	// Require to run in Google App Engine.
	//  Will refactor once I get my docker image working
	err := http.ListenAndServe(":8001", r)
	if err != nil {
		panic(err)
	}

}
