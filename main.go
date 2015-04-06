package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
    "gopkg.in/flosch/pongo2.v3"
	"net/http"
    "fmt"
)

func init() {
	//Config
	//Set Default values
	viper.SetDefault("port", "8080")
	viper.SetDefault("proto", "http")
	viper.SetDefault("staticDir", "./static/")
	viper.SetDefault("templateDir", "./templates/")
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

    //Template setup
    err := pongo2.DefaultSet.SetBaseDirectory(viper.GetString("templateDir"))
    if err != nil {
        fmt.Println("Error setting template base dir:", err.Error())
    }
}

func main() {
	//Create the router
	rtr := mux.NewRouter()
	rtr.StrictSlash(true)
	//Connect our routes in with the router
	connectRoutes(rtr)

	//Hook the router up to the server
	http.Handle("/", rtr)

	//Serve it to the world!
	port := viper.GetString("port")
	http.ListenAndServe(":"+port, nil)

}
