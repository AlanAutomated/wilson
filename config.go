package wilson

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/xeipuuv/gojsonschema"
)

// Convert a JSON configuration to a Configuration struct
func decodeConfig(encodedConfig []byte) (Configuration, error) {
	var config Configuration

	err := json.Unmarshal(encodedConfig, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// Download a JSON configuration and return the body
func downloadConfig(url string) ([]byte, error) {

	result, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, err
	}

	return body, nil

}

// A generic function to read a local configuration file
func readFile(path string) ([]byte, error) {

	config, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Ensure the JSON confinguration conforms to the known schema
func validateConfig(config []byte, schema string) (bool, error) {

	jsonDocument := gojsonschema.NewBytesLoader(config)
	jsonSchema := gojsonschema.NewStringLoader(schema)

	result, err := gojsonschema.Validate(jsonDocument, jsonSchema)

	if err != nil {
		return false, err
	}

	return result.Valid(), nil
}

// A generic function to write a local configuration file
func writeFile(config []byte, path string) error {

	err := ioutil.WriteFile(path, config, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Refresh the configuration on a given interval using a go routine
func RefreshConfig(seconds float64, url string) Configuration {
	time.Sleep(time.Duration(seconds) * time.Second)
	config := Config(url)

	return config
}

// A public function to perform most of the lifting
func Config(url string) Configuration {

	// Track whether a configuration has been read from disk
	fromFile := false

	encodedConfig, err := downloadConfig(url)
	if err != nil {

		// Read from the default configuration path since we don't have a config
		// to tell us where to look
		encodedConfig, err = readFile(".wilson")

		if err != nil {
			log.Fatal(ErrConfigNotFound)
		}

		fromFile = true
	}

	// First check for an error validating the configuration
	valid, err := validateConfig(encodedConfig, Schema)
	if err != nil {
		log.Fatal(ErrConfigNotValid)
	}
	// Next, ensure gojsonschema validated the configuration
	if !valid {
		log.Fatal(ErrConfigNotValid)
	}

	// Decode the configuration from JSON
	config, err := decodeConfig(encodedConfig)
	if err != nil {
		log.Fatal(ErrConfigDecodeFailed)
	}

	// If the file was downloaded from a URL, write it to disk for safe keeping
	if !fromFile {

		path := config.ConfigFile

		err := writeFile(encodedConfig, path)
		if err != nil {
			log.Printf(WarnConfigWriteFailed)
		}
	}

	// Finally, return the config and hope we accounted for all known errors
	return config
}
