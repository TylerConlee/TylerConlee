package main

import (
	"fmt"

	"flag"
	"log"
	"net/http"
	"nrmux"
	
	"github.com/newrelic/go-agent"
)

func main() {
	var nrkey string
	var nrappname string
	flag.StringVar(&nrkey, "nrkey", "null", "New Relic License Key")
	flag.StringVar(&nrappname, "nrappname", "TylerConlee.com", "New Relic Application Name")
	flag.Parse()

	config := newrelic.NewConfig(nrappname, nrkey)
	app, err := newrelic.NewApplication(config)
	if nil != err {
		log.Fatal(err.Error())
	}
	r := nrmux.New(app)
	r.HandleFunc("/", HomeHandler)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Hello world - This is a site")
}
