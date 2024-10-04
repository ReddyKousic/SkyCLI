package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	APIKey   string `json:"apiKey"`
	Locality string `json:"locality"`
}

func SaveAPIKey(apiKey string) error {
	config, err := readConfig()
	if err != nil {
		if os.IsNotExist(err) {
			config = Config{}
		} else {
			return fmt.Errorf("error reading config: %w", err)
		}
	}

	config.APIKey = apiKey

	if err := saveConfig(config); err != nil {
		return err
	}

	fmt.Printf("API key saved: %s\n", apiKey)
	return nil
}

func SaveLocality(locality string) error {
	config, err := readConfig()
	if err != nil {
		if os.IsNotExist(err) {
			config = Config{}
		} else {
			return fmt.Errorf("error reading config: %w", err)
		}
	}

	config.Locality = locality

	if err := saveConfig(config); err != nil {
		return err
	}

	fmt.Printf("Locality saved: %s\n", locality)
	return nil
}

func readConfig() (Config, error) {
	var config Config

	data, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

func saveConfig(config Config) error {
	updatedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	err = os.WriteFile("config.json", updatedData, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}
