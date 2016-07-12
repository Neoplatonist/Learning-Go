package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

// Settings pulls json into config
type Settings struct {
	Hostname       string `toml:"hostname"`
	Protocol       string `toml:"protocol"`
	APIPort        string `toml:"api_port"`
	FileServerPort string `toml:"fileserver_port"`
	WebRoot        string `toml:"webroot"`
	IndexTemplate  string `toml:"index_template"`
	IndexTarget    string `toml:"index_target"`
}

var (
	config     Settings
	configFile = flag.String("conf", "config.toml", "path to toml config")
)

func loadConfig(configFile string) error {
	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		log.Fatal("Error loading config file")
		return err
	}
	return nil
}

func main() {
	log.Println("Initializing...")

	flag.Parse()
	err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Could not load config file: %s\n", *configFile)
	}

	fileServer(config.WebRoot, config.FileServerPort)
}

func fileServer(path string, port string) {
	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)

	log.Println("File Server Listening on: localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
