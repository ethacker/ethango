package ethango


import ("net/http"
	"log"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"encoding/json"
	"io/ioutil"
	"googlemaps.github.io/maps"
)

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

type GeoCodingRequest struct {
	address string
	key     string
}

func getPoliceData(w http.ResponseWriter, r *http.Request){
	apiKey := "AIzaSyDzfdrF3ekdtBR0CFTFvz2F1Qg7BimziTY"
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))


	gc := maps.GeocodingRequest{
		Address: "Main Street",
	}

	results,err := c.Geocode(ctx, &gc)

	log.Print(results)

	policeData, err := client.Get("http://www.dallasopendata.com/resource/are8-xahz.json")
	if(err!=nil) {
		log.Print(err.Error())
	}
	var pdata[] Incident

	var bytes[] byte
	bytes,err = ioutil.ReadAll(policeData.Body)

	if(err!=nil){
		log.Print(err.Error())
	}
	json.Unmarshal(bytes,&pdata)

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}