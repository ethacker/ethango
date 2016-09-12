package main

import (
	"net/http"
	"html/template"
	"path"
	"log"
	"os"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)
var tmplt = make(map[string]*template.Template)

type Contact struct {
	firstname    string
	lastname     string
	emailaddress string
	phonenumber  string
	message      string
}

func init() {
	http.HandleFunc("/contact",contactHandler)
	http.HandleFunc("/", templateHandler)
	tmplt["index"] = template.Must(template.ParseFiles("templates/index.html","templates/layout.html"))
	tmplt["/about"] = template.Must(template.ParseFiles("templates/about.html","templates/layout.html"))
	tmplt["/blog"] = template.Must(template.ParseFiles("templates/blog.html","templates/layout.html"))
	tmplt["/contact"] = template.Must(template.ParseFiles("templates/contact.html","templates/layout.html"))
	tmplt["/projects"] = template.Must(template.ParseFiles("templates/projects.html","templates/layout.html"))
	tmplt["/otherstuff"] = template.Must(template.ParseFiles("templates/otherstuff.html","templates/layout.html"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(appengine.InstanceID())
	if r.Method=="GET" {
		tmplt["/contact"].ExecuteTemplate(w,"layout", nil)
	}
	if r.Method=="POST" {
		ctx := appengine.NewContext(r)
		r.ParseForm()
		c:= Contact{
			firstname: r.FormValue("firstname"),
			lastname: r.FormValue("lastname"),
			emailaddress: r.FormValue("emailaddress"),
			phonenumber: r.FormValue("phonenumer"),
			message: r.FormValue("message"),
		}
		datastore.Put(ctx,datastore.NewIncompleteKey(ctx,"contact",nil),&c)
		log.Print("email " + c.emailaddress + " firstname " + c.firstname + " lastname " + c.lastname)
		log.Print(datastore.Done)
	}
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
