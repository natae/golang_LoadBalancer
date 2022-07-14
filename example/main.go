package main

import (
	"fmt"
	"io/ioutil"
	"nataelb"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"gopkg.in/yaml.v2"
)

func main() {

	config := readConfig()

	go func() {
		lb := nataelb.NewLB()
		if err := lb.Start(config); err != nil {
			panic("LB start failed")
		}
	}()

	go func() {
		startWebServer("First", 5001, "check.php")
	}()

	go func() {
		startWebServer("Second", 5002, "healthcheck")
	}()

	go func() {
		startWebServer("Third", 5003, "check")
	}()

	// Wait for exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-quit
}

func startWebServer(name string, port int, healthcheckPath string) {
	mux := http.NewServeMux()

	// Handle for healthcheck
	mux.HandleFunc(fmt.Sprintf("/%s", healthcheckPath), func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	})

	// Handle for API
	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello ", name)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Println(fmt.Sprintf("Web server \"%s\" started!, port: %d", name, port))
	server.ListenAndServe()
}

func readConfig() *nataelb.Config {
	var config nataelb.Config

	file, err := filepath.Abs("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &config
}
