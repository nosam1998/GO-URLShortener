package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

type ShortUrlPostData struct {
	ShortUrlSlug string
	OriginalUrl  string
}

type JsonResponse struct {
	Success bool
	Message string
}

func checkUrl(urlToTest string) bool {
	_, err := url.ParseRequestURI(urlToTest)
	if err != nil {
		return false
	}

	return true
}

func RedirectRoute(w http.ResponseWriter, r *http.Request) {
	shortUrl := mux.Vars(r)["url"]

	// TODO: Validate shortUrl before query to DB

	originalUrl := GetUrlFromDB(shortUrl)
	// Redirect to originalUrl
	log.Println(originalUrl)
	http.Redirect(w, r, originalUrl, http.StatusPermanentRedirect)
	return
}

func ShortenRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data ShortUrlPostData
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	log.Println(data)

	if checkUrl(data.OriginalUrl) {
		insertUrl(data.OriginalUrl, data.ShortUrlSlug)

		// TODO: Json Response to let the user know that the URL was added to the DB
		log.Println("Json Response to let the user know that the URL was added to the DB")
	} else {
		// TODO: Json Response to let the user know that the URL is invalid
		log.Println("Json Response to let the user know that the URL is invalid")
	}
}
