package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Status - Our struct for status
type Status struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// InstanceIdentity - Instance identity record
type InstanceIdentity struct {
	Codes        string `json:"devpayProductCodes"`
	PrivateIP    string `json:"privateIp"`
	AZ           string `json:"availabilityZone"`
	Version      string `json:"version"`
	Region       string `json:"region"`
	ID           string `json:"instanceId"`
	Billing      string `json:"billingProducts"`
	Type         string `json:"instanceType"`
	ImageID      string `json:"imageId"`
	AccountID    string `json:"accountId"`
	KernelID     string `json:"kernelId"`
	RamdiskID    string `json:"ramdiskId"`
	Architecture string `json:"architecture"`
	PendingTime  string `json:"pendingTime"`
}

func getInstanceIdentity() InstanceIdentity {

	url := "http://169.254.169.254/latest/dynamic/instance-identity/document"

	// Fill the record with the data from the JSON
	var record InstanceIdentity

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return record
	}

	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return record
	}

	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: home")
	fmt.Fprintf(w, "Go-Test!\r\n")
}

func returnHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnHeaders")

	fmt.Println(r.Host)

	w.Header().Set("Cache-Control", "max-age=0")

	// Headers map[string]string `json:"headers"`
	header, _ := json.Marshal(r.Header)
	fmt.Fprintf(w, string(header))
}

func returnStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnStatus")

	w.Header().Set("Cache-Control", "max-age=10")
	status := Status{Status: "OK", Timestamp: time.Now().Format(time.RFC850)}
	json.NewEncoder(w).Encode(status)
}

func returnInstance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnInstance")

	w.Header().Set("Cache-Control", "max-age=60")
	instance := getInstanceIdentity()

	json.NewEncoder(w).Encode(instance)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/headers", returnHeader)
	myRouter.HandleFunc("/status", returnStatus)
	myRouter.HandleFunc("/instance", returnInstance)

	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, myRouter))
}

func main() {
	fmt.Println("Go-Test!")
	fmt.Println("PORT:", os.Getenv("PORT"))

	handleRequests()
}
