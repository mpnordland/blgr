package main

import (
	"github.com/gorilla/mux"
	"net/http"
    "github.com/spf13/viper"
    _ "github.com/mattn/go-sqlite3"
    "github.com/astaxie/beego/orm"
)

func init() {
    //Config
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

    //Database
    orm.RegisterDataBase("default", "sqlite3", "./blgr.db") 
}

func main() {
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
