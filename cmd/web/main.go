package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "8080", "webserver port")
)

func main() {
	flag.Parse()

}

func home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Sever Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, "Internal server Error", 500)
	}
}
