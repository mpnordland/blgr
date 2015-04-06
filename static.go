package main

import (
	"github.com/spf13/viper"
    "github.com/gorilla/mux"
	"net/http"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	staticDir := viper.GetString("staticDir") 
    http.ServeFile(w, r, staticDir+vars["file"])
}
