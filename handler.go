package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := "Pass a name in the query string for a personalized response.\n"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s. Verb called %s.\n", name, r.Method)
	}
	fmt.Fprint(w, message)
}

func addItems(w http.ResponseWriter, r *http.Request) {
	message := "Please pass two valid numbers i.e. val1 and val2\n"
	val1 := r.URL.Query().Get("val1")
	val2 := r.URL.Query().Get("val2")
	iVal1, err := strconv.Atoi(r.URL.Query().Get("val1"))
	iVal2, err := strconv.Atoi(r.URL.Query().Get("val2"))
	if val1 != "" && val2 != "" && err == nil {
		message = fmt.Sprintf("Parameters passed: val1=%s, val2=%s, Result=%s - HTTP Verb called %s\n", val1, val2, strconv.Itoa(iVal1+iVal2), r.Method)
	}
	fmt.Fprint(w, message)
}

func subtractItems(w http.ResponseWriter, r *http.Request) {
	message := "Please pass two valid numbers i.e. val1 and val2\n"
	val1 := r.URL.Query().Get("val1")
	val2 := r.URL.Query().Get("val2")
	iVal1, err := strconv.Atoi(r.URL.Query().Get("val1"))
	iVal2, err := strconv.Atoi(r.URL.Query().Get("val2"))
	if val1 != "" && val2 != "" && err == nil {
		message = fmt.Sprintf("Parameters passed: val1=%s, val2=%s, Result=%s - HTTP Verb called %s\n", val1, val2, strconv.Itoa(iVal1-iVal2), r.Method)
	}
	fmt.Fprint(w, message)
}

func multiplyItems(w http.ResponseWriter, r *http.Request) {
	message := "Please pass two valid numbers i.e. val1 and val2\n"
	val1 := r.URL.Query().Get("val1")
	val2 := r.URL.Query().Get("val2")
	iVal1, err := strconv.Atoi(r.URL.Query().Get("val1"))
	iVal2, err := strconv.Atoi(r.URL.Query().Get("val2"))
	if val1 != "" && val2 != "" && err == nil {
		message = fmt.Sprintf("Parameters passed: val1=%s, val2=%s, Result=%s - HTTP Verb called %s\n", val1, val2, strconv.Itoa(iVal1*iVal2), r.Method)
	}
	fmt.Fprint(w, message)
}

func divideItems(w http.ResponseWriter, r *http.Request) {
	message := "Please pass two valid numbers i.e. val1 and val2\n"
	val1 := r.URL.Query().Get("val1")
	val2 := r.URL.Query().Get("val2")
	iVal1, err := strconv.ParseFloat(r.URL.Query().Get("val1"), 64)
	iVal2, err := strconv.ParseFloat(r.URL.Query().Get("val2"), 64)
	if val1 != "" && val2 != "" && err == nil {
		message = fmt.Sprintf("Parameters passed: val1=%s, val2=%s, Result=%f - HTTP Verb called %s\n", val1, val2, iVal1/iVal2, r.Method)
	}
	fmt.Fprint(w, message)
}

type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

func queueTriggerHandler(w http.ResponseWriter, r *http.Request) {

	var invokeRequest InvokeRequest

	d := json.NewDecoder(r.Body)
	d.Decode(&invokeRequest)

	item := invokeRequest.Data["myQueueItem"]

	outputs := map[string]interface{}{"": ""}

	invokeResponse := InvokeResponse{outputs, nil, ""}

	invokeResponse.Logs = append(invokeResponse.Logs, "hello from queue trigger")
	invokeResponse.Logs = append(invokeResponse.Logs, string(item))

	responseJson, _ := json.Marshal(invokeResponse)
	fmt.Println(string(responseJson))
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/HelloWorld", helloWorld)
	http.HandleFunc("/api/AddNumbers", addItems)
	http.HandleFunc("/api/SubtractNumbers", subtractItems)
	http.HandleFunc("/api/MultiplyNumbers", multiplyItems)
	http.HandleFunc("/api/DivideNumbers", divideItems)
	http.HandleFunc("/queueTrigger", queueTriggerHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
