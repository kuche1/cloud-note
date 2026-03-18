package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/kuche1/cloud-note/lib"
	"github.com/pelletier/go-toml/v2"
)

type Settings struct {
	ServerAddr string
}

func (self Settings) NewFromDefaults() *Settings {
	return &Settings{
		ServerAddr: ":4242",
	}
}

func (self Settings) NewFromPersistentStorage(appStorage fyne.Storage) (*Settings, error) {
	settings := Settings{}.NewFromDefaults()

	settingsFile, _ := getSettingsFile(appStorage)

	data, err := os.ReadFile(settingsFile)
	if err == nil {
		decoder := toml.NewDecoder(strings.NewReader(string(data)))
		// decoder = decoder.DisallowUnknownFields()

		err = decoder.Decode(settings)
		if err != nil {
			return nil, fmt.Errorf("Could not decode settings file: %v", err)
		}
	} else {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("Could not load settings file: %v", err)
		}
	}

	return settings, nil
}

func (self *Settings) Save(appStorage fyne.Storage) error {
	data, err := toml.Marshal(self)
	if err != nil {
		return fmt.Errorf("Could not marshal settings:\n%v", err)
	}

	settingsFile, settingsFileTmp := getSettingsFile(appStorage)

	err = lib.WriteFileAtomic(settingsFile, data, settingsFileTmp)
	if err != nil {
		return fmt.Errorf("Could not save settings to persistent storage:\n%v", err)
	}

	return nil
}

func getSettingsFile(appStorage fyne.Storage) (_realFile string, _temporaryFile string) {
	root := appStorage.RootURI().Path()
	// On PC this is: ~/fyne/could-note
	// This also works on Android

	path := filepath.Join(root, "settings.toml")
	pathTmp := path + ".tmp"

	return path, pathTmp
}
