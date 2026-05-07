package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type SettingValues struct {
	TabString   string   `json:"tabString"`
	DateFormats []string `json:"dateFormats"`
	TimeFormats []string `json:"timeFormats"`
}

var sett *SettingValues

func Settings() *SettingValues {
	if sett == nil {
		var err error
		sett, err = getSettingsFromFile()
		if err != nil {
			sett = defaultSettings()
		}
	}

	return sett
}

func defaultSettings() *SettingValues {
	return &SettingValues{
		TabString: "\t",
		DateFormats: []string{
			"2006-01-02",
			"2006-1-2",
			"01-02-2006",
			"1-2-2006",
			"2006/01/02",
			"2006/1/2",
			"01/02/2006",
			"1/2/2006",
		},
		TimeFormats: []string{
			"15:04:05",
			"03:04:05",
			"3:04:05",
			"3:04:05 PM",
		},
	}
}

func getSettingsFromFile() (*SettingValues, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("Error getting home directory: %w", err)
	}

	settingsPath := filepath.Join(homeDir, ".config", "pasty", "settings.json")

	if _, err := os.Stat(settingsPath); err != nil {
		return nil, fmt.Errorf("Error - missing settings.json: %w", err)
	}

	settingsFile, _ := os.ReadFile(settingsPath)
	var data SettingValues
	err = json.Unmarshal(settingsFile, &data)
	if err != nil {
		return nil, fmt.Errorf("Error reading settings file: %w", err)
	}

	return &data, nil
}
