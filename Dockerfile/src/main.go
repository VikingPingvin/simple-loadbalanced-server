package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/index.html", indexHandler()).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":80", r))
}

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := getContainerId()

		fmt.Fprintf(os.Stderr, "[INFO] Recieved request:\n%v\n", r)

		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		} else {
			w.Write(data)
		}
	})
}
func getContainerId() ([]byte, error) {

	data, err := ioutil.ReadFile("/etc/hostname")

	if err != nil {
		return nil, err
	}

	return data, nil
}
