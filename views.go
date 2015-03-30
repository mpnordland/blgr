package main

import (
    "github.com/astaxie/beego/orm"
    "github.com/gorilla/schema"
    "gopkg.in/flosch/pongo2.v2"
    "net/http"
    "fmt"
)

var decoder = schema.NewDecoder()

func FrontView(w http.ResponseWriter, r *http.Request) {
    o := orm.NewOrm()
    count, err :=o.QueryTable("post").Count()
    fmt.Fprintf(w, "<h1>Hello World!</h1>")
    if err == nil {
        fmt.Fprintf(w, "<p> there are %d posts in the db!</p>", count)
    } else {
        fmt.Fprintf(w, "<p> Error getting number of posts</p>")
    }
}

func AddPostView(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "HEllo!")
    o := orm.NewOrm()
    switch (r.Method){
        case "GET":
            fmt.Fprintf(w, "<p>Still working on actual views, just wanted to get stuff hooked up</p>")
        case "POST":
            err := r.ParseForm()
            if err != nil {
                fmt.Fprintf(w, "<p>Error parsing the form data, try again please!</p>")
                return
            }
            post := new(Post)
            err = decoder.Decode(post, r.PostForm)
            if err != nil {
                fmt.Fprintf(w, "<p>Error parsing the form data, try again please!</p>")
                return
            }
            id, err := o.Insert(post)
            if err != nil {
                fmt.Fprintf(w, "<p>Error saving the form data, try again please!</p>")
                return
            }
            fmt.Fprintf(w, "<p>Added post %d to database!</p>", id)
    }
}
