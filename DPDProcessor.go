package ethango

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
//	"encoding/json"
	"log"
	"golang.org/x/net/context"
	"encoding/json"
	"google.golang.org/appengine/datastore"
	"strings"
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
type Results struct {
	Results []Result `json:"results"`
	Status string `json:"status"`
}

type Result struct {
	AddressComponents []Address `json:"address_components"`
	FormattedAddress string `json:"formatted_address"`
	Geometry Geometry `json:"geometry"`
	PlaceId string `json:"place_id"`
	Types []string `json:"types"`
}

type Address struct {
	LongName string `json:"long_name"`
	ShortName string `json:"short_name"`
	Types []string `json:"types"`
}

type Geometry struct {
	Bounds Bounds `json:"bounds"`
	Location LatLng `json:"location"`
	LocationType string `json:"location_type"`
	Viewport Bounds `json:"viewport"`
}

type Bounds struct {
	Northeast LatLng `json:"northeast"`
	Southwest LatLng `json:"southwest"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
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

func geoCode(w http.ResponseWriter, r *http.Request) {
	c := urlfetch.Client(r)

	var reqBytes [] byte
	var reqLocs [] string

	reqBytes,_ = ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBytes,&reqLocs)

	locations :=make([]LatLng,len(reqLocs))
	for i,v := range reqLocs {
		strings.Replace(v," ","+",-1)
		results,err := c.Get("https://maps.googleapis.com/maps/api/geocode/json?address="+v+"&key="+apiKey)
		if(err!=nil){
			log.Print("geocoding request error")
		}

		var geoResults Results
		var geoBytes [] byte

		geoBytes,_ = ioutil.ReadAll(results)
		json.Unmarshal(geoBytes,&geoResults)

		if(len(geoResults.Results>0)){
			locations[i] = LatLng{
				Lng: geoResults.Results[0].Geometry.Location.Lng,
				Lat: geoResults.Results[0].Geometry.Location.Lat,
			}
		}
	}
}