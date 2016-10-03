package ethango

import (
	"reflect"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type stringTest struct{
	Same bool
	Time time.Duration
}

func mapString(stg string) map[rune]int{
	letterMap := make(map[rune]int)
	for _,v := range stg{
		if(letterMap[v]!=0){
			letterMap[v]++
		} else {
			letterMap[v]=1;
		}
	}
	return  letterMap
}

func permutationHandler(w http.ResponseWriter, r *http.Request){
	var stringArray []string

	bytes,_ :=ioutil.ReadAll(r.Body)
	json.Unmarshal(bytes,&stringArray)
	start := time.Now()
	testResults := permCompare(stringArray[0],stringArray[1])
	elapsed := time.Since(start)
	compareResponse := stringTest{
		Same: testResults,
		Time: elapsed,
	}
	responseBytes,_ := json.Marshal(compareResponse)
	w.Write(responseBytes)
}

func permCompare(first string, second string) bool{
	firstMap := mapString(first)
	secondMap := mapString(second)

	return reflect.DeepEqual(firstMap,secondMap)
}