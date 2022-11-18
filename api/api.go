package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

// ReadAPI takes a URL string argument and a container to store decoded YAML data, and return an error if response code is not 200.
func GetDecodeYAML(url string, div any) error {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("http-get:", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	dec := yaml.NewDecoder(res.Body)
	err = dec.Decode(div)
	if err != nil {
		log.Fatalln("decode:", err)
	}

	return nil
}

// ReadAPI takes a URL string argument and a container to store decoded JSON data, and return an error if response code is not 200.
func GetDecodeJSON(url string, div any) error {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("http-get:", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(div)
	if err != nil {
		log.Fatalln("decode:", err)
	}

	return nil
}
