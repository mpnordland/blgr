package posts 

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"gopkg.in/flosch/pongo2.v3"
	"net/http"
	"strconv"
)

var decoder = schema.NewDecoder()



func PostIndexView(w http.ResponseWriter, r *http.Request) {
    var posts []Post
	o := orm.NewOrm()
	_, err := o.QueryTable("post").All(&posts)
	if err == nil {
		t, err := pongo2.FromFile("posts/index.html")
		if err != nil {
			fmt.Fprintf(w, "<p>Guru Meditation: %s</p>", err.Error())
		} else {
            err = t.ExecuteWriter(pongo2.Context{"posts": posts}, w)
            if err != nil {
                fmt.Fprintf(w, "<p>Guru Meditation: %s</p>", err.Error())
            }
        }
    } else {
		fmt.Fprintf(w, "<p> Error getting number of posts</p>")
	}
}

func AddPostView(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()
	switch r.Method {
	case "GET":
		t, err := pongo2.FromFile("posts/create.html")
		if err != nil {
			fmt.Fprintf(w, "<p>Still working on actual views, just wanted to get stuff hooked up</p>")
			fmt.Fprintf(w, "<p>Guru Meditation: %s</p>", err.Error())
		} else {
			err = t.ExecuteWriter(pongo2.Context{}, w)
			if err != nil {
				fmt.Fprintf(w, "<p>Houston, we have a problem.</p>")
			}
		}
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
		if err != nil { fmt.Fprintf(w, "<p>Error saving the form data, try again please!</p>")
			fmt.Fprintf(w, "<p>Guru Meditation: %s</p>", err.Error())
			return
		}
		fmt.Fprintf(w, "<p>Added post %d to database!</p>", id)
	}
}

func PostView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	o := orm.NewOrm()
	postId, err := strconv.Atoi(vars["post"])
	if err != nil {
		fmt.Fprintf(w, "<p>Guru Meditation: %s</p>", err.Error())
		return
	}
	post := Post{Id: postId}
	err = o.Read(&post)
	if err != nil {
		fmt.Fprintf(w, "<p>Guru Meditation: %s</p>", err.Error())
		return
	}
	tpl, err := pongo2.FromFile("posts/view.html")
	if err != nil {
		fmt.Fprintf(w, "<p>Guru Meditation: %s</p>", err.Error())
	} else {
		err := tpl.ExecuteWriter(pongo2.Context{"post": post}, w)
		if err != nil {
			fmt.Fprintf(w, "<p>Guru Meditation: %s</p>", err.Error())
		}
	}
}
