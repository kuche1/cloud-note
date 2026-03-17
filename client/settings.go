package client

import (
	"os"

	"fyne.io/fyne/v2/storage"
)

// IMPROVE: Actually make a struct with some settings
func LoadServerAddr() (_addr string, _alreadySet bool, _err error) {
	settingsFile := getConfigFile()

	data, err := os.ReadFile(settingsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", false, nil
		}
		return "", false, err
	}

	return string(data), true, nil
}

func SaveServerAddr(addr string) error {
	settingsFile := getConfigFile()

	err := os.WriteFile(settingsFile, []byte(addr), 0600)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFile() string {
	uri := storage.NewFileURI("myfile.txt")
	return uri.Path()

	// TODO: `os.UserConfigDir` does not work on android
	// val, err := os.UserConfigDir()
	// if err != nil {
	// 	return "", fmt.Errorf("Could not get configuration file:\n%v", err)
	// }
	// // IMPROVE: Actually create a folder, do not use a file
	// return filepath.Join(val, "could-note-server-addr.txt"), nil
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
