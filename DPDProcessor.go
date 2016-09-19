package ethango


import ("net/http"
	"log"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"encoding/json"
	"io/ioutil"
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

func getPoliceData(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
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