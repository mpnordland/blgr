package posts

import "github.com/gorilla/mux"

func SetRoutes(r *mux.Router) {
    r.HandleFunc("/", PostIndexView)
	r.HandleFunc("/add/", AddPostView)
	r.HandleFunc("/{post:[0-9]+}/", PostView)
}
