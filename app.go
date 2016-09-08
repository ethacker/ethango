package hello

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"./views/")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"./views/about.html")
}
