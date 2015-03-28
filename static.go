package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/spf13/viper"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    staticDir := viper.GetString("staticDir")
    http.ServeFile(w, r, staticDir+vars["file"])
}
