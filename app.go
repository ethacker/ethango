package hello

import (
	"net/http"
	"html/template"
	"path"
	"log"
	"os"
)

//var templates = template.Must(template.ParseGlob("templates/*"))

var tmplt = make(map[string]*template.Template)

func init() {
	http.HandleFunc("/", templateHandler)
	tmplt["index"] = template.Must(template.ParseFiles("templates/index.html","templates/layout.html"))
	tmplt["/about"] = template.Must(template.ParseFiles("templates/about.html","templates/layout.html"))
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)
	if r.URL.Path != "/"{
		fp := path.Join("templates", r.URL.Path + ".html")

		info, err := os.Stat(fp)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
		}

		if info.IsDir() {
			http.NotFound(w, r)
			return
		}

		if err != nil {
			// Log the detailed error
			log.Println(err.Error())
			// Return a generic "Internal Server Error" message
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if err := tmplt[r.URL.Path].ExecuteTemplate(w, "layout", nil); err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
		}


	} else {

		tmplt["index"].ExecuteTemplate(w,"layout", nil)
	}

}
