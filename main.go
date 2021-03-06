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
	s := r.PathPrefix("/api/").Subrouter()
	s.HandleFunc("/episodes", middleware.HandleFunc(AllEps))
	s.HandleFunc("/episode", middleware.HandleFunc(AllEps))
	s.HandleFunc("/episode/{id}", middleware.HandleFunc(SingleEpisode))
	s.HandleFunc("/stats", middleware.HandleFunc(StatsIndex))
	s.HandleFunc("/", middleware.HandleFunc(Index))
	fmt.Println("Routes set up")
	// Catches anything left
	http.Handle("/", r)

	// Require to run in Google App Engine.
	//  Will refactor once I get my docker image working
	err := http.ListenAndServe(":8001", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 8001")

}
