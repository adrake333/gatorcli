package config




import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)




const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL		string `json:"db_url"`
	CurrentUserName	string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)

	return fullPath, nil
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	body, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}
	var config Config

	err = json.Unmarshal(body, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	write(c)
	return nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
