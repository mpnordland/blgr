package main

import (
    "github.com/gorilla/mux"
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
    postRoutes(postR)

    //Serve static files
    router.Handle("/static/{file}", http.HandlerFunc(staticHandler))

}

func feedRoutes(r *mux.Router) {
}

func postRoutes(r *mux.Router) {
    r.HandleFunc("/add/", AddPostView)    
}

