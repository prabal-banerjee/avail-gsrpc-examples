package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Seed   string `json:"seed"`
	ApiURL string `json:"api_url"`
}

func (c *Config) GetConfig(configFileName string) error {

	jsonFile, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return err
	}

	return nil
}
