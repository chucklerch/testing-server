package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port     string = ":8080"
	bgColor  string = "white"
	version  string = "v1.0.0" // Displayed version.
	hostname string = ""
	podname  string = ""
	//tls     bool = false
)

func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html><html><head><title>Homepage</title></head>\n")
	fmt.Fprintf(w, "<body style=\"background-color:%s;\">", bgColor)
	fmt.Fprintf(w, "<p>Version %s\n", version)
	fmt.Fprintf(w, "<p>Hostname %s\n", hostname)
	fmt.Fprintf(w, "<p>Pod Name %s\n", podname)
	fmt.Fprintf(w, "<h4>Links</h4>")
	fmt.Fprintf(w, "<p><a href=\"/headers\">Headers</a>\n")
	fmt.Fprintf(w, "<p><a href=\"/status?code=200\">Status</a>\n")
	fmt.Fprintf(w, "</body></html>\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func statusCode(w http.ResponseWriter, req *http.Request) {
	var code string
	val, ok := req.URL.Query()["code"]
	if !ok || len(val[0]) < 1 {
		code = "200"
	} else {
		code = val[0]
	}

	i, err := strconv.Atoi(code)
	if err != nil {
		i = 500
	}
	w.WriteHeader(i)
	fmt.Fprintf(w, "code: %v\n", code)
}

func main() {
	val, ok := os.LookupEnv("PORT")
	if ok {
		port = ":" + val
	}
	val, ok = os.LookupEnv(("BGCOLOR"))
	if ok {
		bgColor = val
	}
	val, ok = os.LookupEnv(("VERSION"))
	if ok {
		version = val
	}
	val, ok = os.LookupEnv(("HOSTNAME"))
	if ok {
		hostname = val
	}
	val, ok = os.LookupEnv(("PODNAME"))
	if ok {
		podname = val
	}
	// Setup handlers
	http.HandleFunc("/", homePage)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/status", statusCode)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
	//   log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
