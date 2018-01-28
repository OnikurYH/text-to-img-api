package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Font struct {
		Include  map[string]bool
		Defaults struct {
			Font        string
			FontSize    float64 `yaml:"fontSize"`
			MinFontSize float64 `yaml:"minFontSize"`
			MaxFontSize float64 `yaml:"maxFontSize"`
			FontColor   string  `yaml:"fontColor"`
		}
	}
}

var config Config

// Get ...
func Get() Config {
	return config
}

// Init ...
func Init() {
	dat, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(dat, &config); err != nil {
		log.Fatal(err)
	}
}
