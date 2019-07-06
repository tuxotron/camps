package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/thedevsaddam/gojsonq"
)

var campsBytes []byte

func campsEndPoint(w http.ResponseWriter, r *http.Request) {

	jq := gojsonq.New().FromString(string(campsBytes))

	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
	}

	state := r.Form.Get("state")

	var res interface{}
	if len(state) > 0 {
		res = jq.From("camps").Where("state", "=", state).Get()
	} else {
		res = jq.From("camps").Get()
	}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func handleRequests() {
	http.HandleFunc("/camps/", campsEndPoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadCamps() {
	jsonFile, err := os.Open("camps.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	campsBytes, _ = ioutil.ReadAll(jsonFile)
}

func main() {
	loadCamps()
	handleRequests()
}
