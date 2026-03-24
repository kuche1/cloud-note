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
	// private (fields with lower case letters do not get saved to the settings file)
	persistentStorage string

	// actual settings
	ServerAddr     string
	ServerPassword string

	// app-related persistent stuff
	LastEditedNote string
}

func (self Settings) NewFromDefaults(persistentStorage string) *Settings {
	return &Settings{
		persistentStorage: persistentStorage,
		ServerAddr:        "",
		ServerPassword:    "",
		LastEditedNote:    "",
	}
}

func (self *Settings) LoadFromPersistentStorage() error {
	settingsFile, _ := getSettingsFile(self.persistentStorage)

	data, err := os.ReadFile(settingsFile)
	if err == nil {
		decoder := toml.NewDecoder(strings.NewReader(string(data)))
		// decoder = decoder.DisallowUnknownFields()

		err = decoder.Decode(self)
		if err != nil {
			return fmt.Errorf("Could not decode settings file:\n%v", err)
		}
	} else {
		if os.IsNotExist(err) {
			err := self.Save()
			if err != nil {
				return fmt.Errorf("Could not save settings:\n%v", err)
			}
		} else {
			return fmt.Errorf("Could not load settings file:\n%v", err)
		}
	}

	return nil
}

func (self *Settings) Save() error {
	data, err := toml.Marshal(self)
	if err != nil {
		return fmt.Errorf("Could not marshal settings:\n%v", err)
	}

	settingsFile, settingsFileTmp := getSettingsFile(self.persistentStorage)

	err = lib.FileWriteAtomic(settingsFile, data, settingsFileTmp)
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
