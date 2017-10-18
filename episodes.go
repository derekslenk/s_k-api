package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Episode contains all the information for an episode object
type Episode struct {
	EpisodeNumber int    `json:"epNumber"`
	ReleaseDate   string `json:"releaseDate"`
	Special       bool   `json:"special"`
	Desc          string `json:"desc"`
	Duration      int    `json:"duration_seconds"`
	Libsyn        string `json:"libsyn"`
}

// Episodes are an array of type Episode
type Episodes []Episode

// SingleEpisode returns episode number {id}
//  from the database
func SingleEpisode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: SingleEpisode")
	vars := mux.Vars(r)
	key := vars["id"]

	rows := db.QueryRow("SELECT epNumber, releaseDate, special, description, duration_seconds, libsyn FROM episodes WHERE epNumber=?", key)
	ep := new(Episode)
	err := rows.Scan(&ep.EpisodeNumber, &ep.ReleaseDate, &ep.Special, &ep.Desc, &ep.Duration, &ep.Libsyn)

	if err != nil {
		//json.NewEncoder(w).Encode("could not find row " + key)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("could not find row " + key)
	} else {
		json.NewEncoder(w).Encode(ep)
		fmt.Printf("Key: " + key)
	}

}

// AllEps returns all the episodes
func AllEps(w http.ResponseWriter, r *http.Request) {
	// Coming soon, database!
	rows, err := db.Query("SELECT epNumber, releaseDate, special, description, duration_seconds, libsyn FROM episodes")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("couldn't select from the table")
	} else {
		defer rows.Close()

		episodes := make([]*Episode, 0)
		for rows.Next() {
			ep := new(Episode)
			err := rows.Scan(&ep.EpisodeNumber, &ep.ReleaseDate, &ep.Special, &ep.Desc, &ep.Duration, &ep.Libsyn)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode("couldn't read the row")
			}
			episodes = append(episodes, ep)
		}
		if err = rows.Err(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("something fucked up")
		}

		fmt.Println("Endpoint Hit: All Episodes")

		json.NewEncoder(w).Encode(episodes)

	}

}

// EpCount counts the number of episodes
func EpCount() (count int) {
	rows, err := db.Query("SELECT COUNT(*) FROM episodes")
	if err != nil {
		panic(err)
	} else {
		// This is really important for some reason
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&count)
			if err != nil {
				panic(err)
			}
		}
	}

	return count
}

// SpecialEpCount counts the number of episodes
func SpecialEpCount() (count int) {
	rows, err := db.Query("SELECT COUNT(*) FROM episodes WHERE special = true")
	if err != nil {
		panic(err)
	} else {
		// This is really important for some reason
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&count)
			if err != nil {
				panic(err)
			}
		}
	}

	return count
}

// AvgLength returns the average length over all episodes (in seconds)
func AvgLength() (length int) {
	rows, err := db.Query("SELECT ROUND((AVG(duration_seconds) / 60) ,0)  FROM episodes WHERE special = false")
	if err != nil {
		panic(err)
	} else {
		// This is really important for some reason
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&length)
			if err != nil {
				panic(err)
			}
		}
	}

	return length
}

// AvgLengthSpecial returns the average length over all episodes (in seconds)
func AvgLengthSpecial() (length int) {
	rows, err := db.Query("SELECT ROUND((AVG(duration_seconds) / 60) ,0) FROM episodes")
	if err != nil {
		panic(err)
	} else {
		// This is really important for some reason
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&length)
			if err != nil {
				panic(err)
			}
		}
	}

	return length
}

// InsertEpisode creates an episode type?
//  but inserts into the DB
func InsertEpisode(epNumber int, published time.Time, special bool, description string, durationSeconds int, link string) (success bool) {
	//fmt.Println("This is where I would insert")
	result, err := db.Exec("INSERT INTO episodes (releaseDate, special, description, duration_seconds, libsyn) VALUES (?,?,?,?,?)", published, special, description, durationSeconds, link)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Inserted episode: ", result)
	return true
}
