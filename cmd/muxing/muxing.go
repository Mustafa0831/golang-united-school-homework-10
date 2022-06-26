package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func name(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)["PARAM"]
	w.Write([]byte(fmt.Sprintf("Hello, %s!", p)))
}

func bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	return
}

func data(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", b)))
}

func headers(w http.ResponseWriter, r *http.Request){
	a:=r.Header.Get("a")
	b:=r.Header.Get("b")

	aAtoi,err:=strconv.Atoi(a)
	if err!= nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bAtoi,err:=strconv.Atoi(b)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("a+b",strconv.Itoa(aAtoi+bAtoi))
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.Path("/name/{PARAM}").HandlerFunc(name).Methods(http.MethodGet)
	router.Path("/bad").HandlerFunc(bad).Methods(http.MethodGet)
	router.Path("/data").HandlerFunc(data).Methods(http.MethodPost)
	router.Path("/headers").HandlerFunc(headers).Methods(http.MethodPost)


	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
