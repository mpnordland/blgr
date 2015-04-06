package main 

import (
	"github.com/astaxie/beego/orm"
	_"gopkg.in/flosch/pongo2.v3"
	"net/http"
	"fmt"
)


func FrontView(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()
	count, err := o.QueryTable("post").Count()
	if err == nil {
        fmt.Fprintf(w, "<h1>Hello World</h1><p>there are %d posts in the db!</p>", count)
    } else {
		fmt.Fprintf(w, "<p> Error getting number of posts</p>")
	}
}

