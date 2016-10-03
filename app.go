package ethango

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
	"google.golang.org/appengine/user"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"google.golang.org/appengine/urlfetch"
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

type Usr struct {
	Email      string
	AuthDomain string
	Admin      bool
	Date 	   time.Time
}

type test struct {number int}

func init() {
	http.HandleFunc("/api/strings",permutationHandler)
	http.HandleFunc("/api/kibo",oracleDistro)
	http.HandleFunc("/dpdinfo/cron/crimes",saveCrimeData)
	http.HandleFunc("/dpdinfo/crimes",getPoliceData)
	http.HandleFunc("/exclusive",exclusiveHandler)
	http.HandleFunc("/contact",contactHandler)
	http.HandleFunc("/", templateHandler)
	tmplt["index"] = template.Must(template.ParseFiles("templates/index.html","templates/layout.html"))
	tmplt["/about"] = template.Must(template.ParseFiles("templates/about.html","templates/layout.html"))
	tmplt["/blog"] = template.Must(template.ParseFiles("templates/blog.html","templates/layout.html"))
	tmplt["/contact"] = template.Must(template.ParseFiles("templates/contact.html","templates/layout.html"))
	tmplt["/projects"] = template.Must(template.ParseFiles("templates/projects.html","templates/layout.html"))
	tmplt["/otherstuff"] = template.Must(template.ParseFiles("templates/otherstuff.html","templates/layout.html"))
	tmplt["/exclusive"] = template.Must(template.ParseFiles("templates/exclusive.html","templates/layout.html"))
	tmplt["/login"] = template.Must(template.ParseFiles("templates/login.html","templates/layout.html"))
	tmplt["/dpdinfo"] = template.Must(template.ParseFiles("templates/dpdinfo.html","templates/layout.html"))
	tmplt["/strings"] = template.Must(template.ParseFiles("templates/strings.html","templates/layout.html"))
}


func userKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c,"User","registered_users",0,nil)
}

func exclusiveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if (u==nil) {
		url,err :=user.LoginURL(ctx, "/exclusive")
		if(err!=nil){
			http.Error(w,err.Error(),http.StatusTeapot)
		}
		fmt.Fprintf(w, `<a href="%s">Sign in to access this Exclusive Itemz</a>`, url)
	} else {
		usr := Usr{
			Email: u.Email,
			AuthDomain: u.AuthDomain,
			Admin: u.Admin,
			Date: time.Now(),
		}
		saveUser(w,r,usr)
		tmplt["/exclusive"].ExecuteTemplate(w,"layout",nil)
	}
}

func saveUser(w http.ResponseWriter, r *http.Request,u Usr) {
	c := appengine.NewContext(r)

	key := datastore.NewIncompleteKey(c,"User",userKey(c))
	_, err:= datastore.Put(c,key,&u)
	if err !=nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET" {
		w.Header().Add("Content-type","text/html")
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
		Firstname: html.EscapeString(r.FormValue("firstname")),
		Lastname: html.EscapeString(r.FormValue("lastname")),
		Emailaddress: html.EscapeString(r.FormValue("emailaddress")),
		Phonenumber: html.EscapeString(r.FormValue("phonenumber")),
		Message: html.EscapeString(r.FormValue("message")),
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
	w.Header().Set("Content-type","text/html")
	user.LoginURL(appengine.NewContext(r),"/exclusive")
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

func oracleDistro(w http.ResponseWriter, r *http.Request) {
	ctx :=appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp,_ := client.Get("https://storage.googleapis.com/oracleimage/Oracle11gHeadless.ova.gz")
	defer resp.Body.Close()


	w.Header().Set("Content-Disposition", "attachment; filename=OracleHeadlessImage.ova.gz")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w,resp.Body)
}