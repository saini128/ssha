package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"ssha/models"
)

const configFileName = "config.json"

func getConfigFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".ssha", configFileName), nil
}

func GetHosts() ([]models.Host, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			// If config file doesn't exist, create it with default values.
			defaultHosts := []models.Host{}
			if err := saveHosts(defaultHosts); err != nil {
				return nil, err
			}
			return defaultHosts, nil // return default hosts after creation.
		}

		return nil, err
	}

	var hosts []models.Host
	if err := json.Unmarshal(data, &hosts); err != nil {
		return nil, err
	}

	return hosts, nil
}

func saveHosts(hosts []models.Host) error {
	configFile, err := getConfigFile()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(hosts, "", "    ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(configFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}
	return os.WriteFile(configFile, jsonData, 0600)
}

func SaveHost(host models.Host) error {
	hosts, err := GetHosts()
	if err != nil {
		return err
	}
	hosts = append(hosts, host)
	return saveHosts(hosts)

}

func UpdateHost(updatedHost models.Host) error {
	hosts, err := GetHosts()
	if err != nil {
		return err
	}

	for i, host := range hosts {
		if host.Alias == updatedHost.Alias {
			hosts[i] = updatedHost
			return saveHosts(hosts)
		}
	}
	return fmt.Errorf("host with alias %s not found", updatedHost.Alias)
}

func DeleteHost(alias string) error {
	hosts, err := GetHosts()
	if err != nil {
		return err
	}

	for i, host := range hosts {
		if host.Alias == alias {
			hosts = append(hosts[:i], hosts[i+1:]...)
			return saveHosts(hosts)
		}
	}
	return fmt.Errorf("host with alias %s not found", alias)
}
