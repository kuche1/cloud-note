package settings

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func (self *Settings) SetServerAddr(appStorage fyne.Storage, new string) error {
	old := self.ServerAddr

	self.ServerAddr = new

	err := self.Save(appStorage)
	if err != nil {
		self.ServerAddr = old
		return fmt.Errorf("Could not set server address:\n%v", err)
	}

	return nil
}
