package client

import (
	"os"
	"path/filepath"
)

// IMPROVE: Actually make a struct with some settings
func LoadServerAddr(app *App) (_addr string, _alreadySet bool, _err error) {
	settingsFile := getConfigFile(app)

	data, err := os.ReadFile(settingsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", false, nil
		}
		return "", false, err
	}

	return string(data), true, nil
}

func SaveServerAddr(app *App, addr string) error {
	settingsFile := getConfigFile(app)

	err := os.WriteFile(settingsFile, []byte(addr), 0600)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFile(app *App) string {
	root := app.app.Storage().RootURI().Path()
	// On PC this is: ~/fyne/could-note
	// This also works on Android

	return filepath.Join(root, "myfile.txt")
}

// type Settings struct {
// 	ServerAddr string
// }

// func (self Settings) NewFromDefaults() *Settings {
// 	return &Settings{
// 		ServerAddr: ":4242",
// 	}
// }

// func (self Settings) NewFromConfig() (*Settings, error) {
// 	settings := Settings{}.NewFromDefaults()

// 	// TODO: Actually use a config located in the config directory
// 	data, err := os.ReadFile("settings.toml")
// 	if err == nil {
// 		decoder := toml.NewDecoder(strings.NewReader(string(data)))
// 		// decoder = decoder.DisallowUnknownFields()

// 		err = decoder.Decode(settings)
// 		if err != nil {
// 			return nil, fmt.Errorf("Could not decode settings file: %v", err)
// 		}
// 	} else {
// 		if !os.IsNotExist(err) {
// 			return nil, fmt.Errorf("Could not load settings file: %v", err)
// 		}
// 	}

// 	return settings, nil
// }
