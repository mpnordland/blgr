package main

import (
	"github.com/gorilla/mux"
	"net/http"
    "github.com/spf13/viper"
)

//Thus named because of Go's reserved
//package init() function. see here:
// http://golang.org/ref/spec#Program_initialization_and_execution
func initialize() {
    //Set Default values
    viper.SetDefault("port", "8080")
    viper.SetDefault("proto", "http")
    viper.SetDefault("staticDir", "./static/")
    viper.SetDefault("templateDir", "./template/")
    //Read in config file
    viper.SetConfigName("config")
    viper.AddConfigPath("/etc/blgr/")
    viper.AddConfigPath("$HOME/.blgr")
    viper.AddConfigPath(".")
    viper.ReadInConfig()
    //Allow flag or environment overrides
    viper.SetEnvPrefix("BLGR")
    viper.AutomaticEnv()
}

func main() {
    //Setup config values
    initialize()
    
    //Create the router
	rtr := mux.NewRouter()

    //Connect our routes in with the router
    connectRoutes(rtr)

    //Hook the router up to the server
	http.Handle("/", rtr)
    
    //Serve it to the world!
    port := viper.GetString("port")
	http.ListenAndServe(":"+port, nil)

}
