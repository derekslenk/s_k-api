package main

import (
	"fmt"
	"net/http"

	"github.com/derekslenk/gomodels"
)

// Route defines our structure for the class
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array of many route
type Routes []Route

// Index just returns some non-json information
//  Should probably be updated
func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Steve and Kyle Podcast: #api")
	fmt.Fprintln(w, "Number of episodes in database:", models.EpCount())
	fmt.Fprintln(w, "Created by Derek Slenk")
	fmt.Println("Endpoint Hit: Index")
}
