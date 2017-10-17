package main

import (
	"net/http"
	"os"

	"github.com/betacraft/yaag/middleware"
	"github.com/betacraft/yaag/yaag"
	"google.golang.org/appengine"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var appPort string
var version string

func init() {
	// Get and open a DB connections
	// EXAMPLE MYSQL_CONNECTION: username:password@tcp(dbserver.com(or IP address):3306)/databaseName
	InitDB(os.Getenv("MYSQL_CONNECTION"))

	// This is for generating API documentation. I am not sure how I feel about it
	yaag.Init(&yaag.Config{On: false, DocTitle: "s_k-api", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "https://steve-and-kyle.appspot.com", "Staging": "iCantAffordThat.sorry"}})

	r := mux.NewRouter()
	r.HandleFunc("/episodes", middleware.HandleFunc(AllEps))
	r.HandleFunc("/episode", middleware.HandleFunc(AllEps))
	r.HandleFunc("/episode/{id}", middleware.HandleFunc(SingleEpisode))
	r.HandleFunc("/stats", middleware.HandleFunc(StatsIndex))
	r.HandleFunc("/", middleware.HandleFunc(Index))

	// Catches anything left
	http.Handle("/", r)
}

func main() {
	// Require to run in Google App Engine.
	//  Will refactor once I get my docker image working
	appengine.Main()

}
