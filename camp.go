package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Camps struct {
	Camps []Camp `json:"camps"`
}

// User struct which contains a name
// a type and a list of social links
type Camp struct {
	Name    string `json:"name"`
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Phone   string `json:"phone"`
	Fax     string `json:"fax"`
}

var camps Camps

func campPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Camp!")
}

func icePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ICE")
}

func handleRequests() {
	http.HandleFunc("/camp", campPage)
	http.HandleFunc("/ice", icePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadCamps() {
	jsonFile, err := os.Open("camps.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &camps)
}

func main() {

	loadCamps()

	for i := 0; i < len(camps.Camps); i++ {
		fmt.Println("Camp Name: " + camps.Camps[i].Name)
	}

	handleRequests()
}
