package main

import (
	"github.com/gorilla/mux"
    "github.com/mpnordland/blgr/posts"
	"net/http"
)

//Here's where we define the routes for
//the application
func connectRoutes(router *mux.Router) {
	//Create subrouters for these paths
	feedR := router.PathPrefix("/feed/").Subrouter()
	postR := router.PathPrefix("/posts").Subrouter()

	//Front page
	router.Handle("/", http.HandlerFunc(FrontView))

	//Connect routes
	feedRoutes(feedR)
	posts.SetRoutes(postR)

	//Serve static files
	router.HandleFunc("/static/{file:.+}", staticHandler)

}

func feedRoutes(r *mux.Router) {
}

