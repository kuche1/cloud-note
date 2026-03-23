package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kuche1/cloud-note/lib"
	"github.com/pelletier/go-toml/v2"
)

type Settings struct {
	// private
	persistentStorage string

	// actual settings
	ServerAddr string
	// ServerPassword string // IMPROVE001: Add a "user token"/"server password" -> if this field is empty, ask the user for a password, then fill in that field with the password, and connect to the server

	// app-related persistent stuff
	LastEditedNote string
}

func (self Settings) NewFromDefaults(persistentStorage string) *Settings {
	return &Settings{
		persistentStorage: persistentStorage,
		ServerAddr:        "localhost:4242",
		// ServerPassword:    "",
		LastEditedNote: "",
	}
}

func (self Settings) NewFromPersistentStorage(persistentStorage string) (*Settings, error) {
	settings := Settings{}.NewFromDefaults(persistentStorage)

	settingsFile, _ := getSettingsFile(settings.persistentStorage)

	data, err := os.ReadFile(settingsFile)
	if err == nil {
		decoder := toml.NewDecoder(strings.NewReader(string(data)))
		// decoder = decoder.DisallowUnknownFields()

		err = decoder.Decode(settings)
		if err != nil {
			return nil, fmt.Errorf("Could not decode settings file:\n%v", err)
		}
	} else {
		if os.IsNotExist(err) {
			err := settings.Save()
			if err != nil {
				return nil, fmt.Errorf("Could not save settings:\n%v", err)
			}
		} else {
			return nil, fmt.Errorf("Could not load settings file:\n%v", err)
		}
	}

	return settings, nil
}

func (self *Settings) Save() error {
	data, err := toml.Marshal(self)
	if err != nil {
		return fmt.Errorf("Could not marshal settings:\n%v", err)
	}

	settingsFile, settingsFileTmp := getSettingsFile(self.persistentStorage)

	err = lib.WriteFileAtomic(settingsFile, data, settingsFileTmp)
	if err != nil {
		return fmt.Errorf("Could not save settings to persistent storage:\n%v", err)
	}

	return nil
}

func getSettingsFile(persistentStorage string) (_realFile string, _temporaryFile string) {
	path := filepath.Join(persistentStorage, "settings.toml")
	pathTmp := path + ".tmp"

	return path, pathTmp
}
