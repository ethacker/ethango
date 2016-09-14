package main

import (
	"net/http"
	"html/template"
	"path"
	"log"
	"os"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
	"time"
	"golang.org/x/net/context"
)
var tmplt = make(map[string]*template.Template)

type Contact struct {
	Firstname    string
	Lastname     string
	Emailaddress string
	Phonenumber  string
	Message      string
	Date	     time.Time
}

type test struct {number int}

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
	if r.Method=="GET" {
		tmplt["/contact"].ExecuteTemplate(w,"layout", nil)
	}
	if r.Method=="POST" {
		r.ParseForm()

		saveContact(w,r)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func contactKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "Contact", "default_contacts", 0, nil)
}

func saveContact(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	ctc := &Contact{
		Firstname: r.FormValue("firstname"),
		Lastname: r.FormValue("lastname"),
		Emailaddress: r.FormValue("emailaddress"),
		Phonenumber: r.FormValue("phonenumber"),
		Message: r.FormValue("message"),
		Date: time.Now(),
	}
	log.Print(ctc.Firstname)

	key := datastore.NewIncompleteKey(c, "Message", contactKey(c))
	_, err := datastore.Put(c, key, ctc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
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