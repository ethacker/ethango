package ethango

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"log"
	"golang.org/x/net/context"
	"encoding/json"
	"google.golang.org/appengine/datastore"
)

var apiKey ="AIzaSyDzfdrF3ekdtBR0CFTFvz2F1Qg7BimziTY"

type Incident struct {
	Beat string `json:"beat"`
	Block string `json:"block"`
	Location string `json:"location"`
	Nature string `json:"nature_of_call"`
	Priority string `json:"priority"`
	Date string `json:"date_time"`
	Division string `json:"division"`
	IncNumber string `json:"incident_number"`
	ReportingArea string `json:"reporting_area"`
	Status string `json:"status"`
	UnitNumber string `json:"unit_number"`
}

func incidentKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c,"Incident","dpd_incidents",0,nil)
}

func saveIncident(w http.ResponseWriter, r *http.Request,i Incident) {
	c := appengine.NewContext(r)

	var incident []Incident

	datastore.NewQuery("Incident").Filter("IncNumber =", i.IncNumber).Limit(1).GetAll(c,&incident)
	if(!(len(incident)>0)){
		log.Print("saving incident" + i.IncNumber)
		key := datastore.NewIncompleteKey(c,"Incident",incidentKey(c))
		_, err:= datastore.Put(c,key,&i)
		if err !=nil {
			log.Print(err.Error())
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
	}

}

func getPoliceData(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)


	policeData, err := client.Get("http://www.dallasopendata.com/resource/are8-xahz.json")
	if(err!=nil) {
		log.Print(err.Error())
	}

	var bytes[] byte
	var incidents[] Incident

	bytes,err = ioutil.ReadAll(policeData.Body)
	json.Unmarshal(bytes,&incidents)

	for _,v := range incidents {
		saveIncident(w,r,v)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}