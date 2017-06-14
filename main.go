package main

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

type BupwConfig struct {
	Httpserver struct {
		Port  string `yaml:"port"`
		Query string `yaml:"query"`
	}
	Files []string
}

type BupwResponse struct {
	Status bool `json:"status"`
}

var bupwConfig *BupwConfig
var words map[string]bool

func main() {
	config, err := getProjectConfig()
	bupwConfig = config

	if err != nil {
		log.Fatalln(err.Error())
	}

	words = readFiles(config.Files)

	log.Printf("Imported words: %d \n", len(words))

	if err != nil {
		log.Fatalln(err.Error())
	}

	http.HandleFunc("/", handleRequest)
	log.Printf("Start server: %s \n", ":"+config.Httpserver.Port)

	http.ListenAndServe(":"+config.Httpserver.Port, nil)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	passwordQuery := r.URL.Query().Get(bupwConfig.Httpserver.Query)

	log.Printf("Incoming password: %s \n", passwordQuery)

	status := false
	if words[passwordQuery] {
		status = true
	}

	// json resposne
	js, err := json.Marshal(BupwResponse{status})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send reponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
}

func getProjectConfig() (*BupwConfig, error) {
	config := &BupwConfig{}

	if err := YamlUnmarshal("./config.yml", config); err != nil {
		return nil, err
	}
	return config, nil
}

func YamlUnmarshal(path string, out interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bytes, out); err != nil {
		return err
	}
	return nil
}
