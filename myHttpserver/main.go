package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")


	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request){
	fmt.Println("entering /healthz handler")
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))

	} else {
		io.WriteString(w, "hello [stranger]\n")

	}

	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		//io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		w.Write([]byte(fmt.Sprintf("%s=%s\n", k, v)))

	}
	w.WriteHeader(200)
	w.Write([]byte(r.RemoteAddr+"\n"))
	version := os.Getenv("VERSION")
	w.Write([]byte(fmt.Sprintf("%s=%s\n", "version", version)))

	glog.V(2).Info("client ip:"+r.Host)


}
