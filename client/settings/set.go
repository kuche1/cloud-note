package settings

import (
	"fmt"
)

func (self *Settings) SetServerAddr(new string) error {
	old := self.ServerAddr

	self.ServerAddr = new

	err := self.Save()
	if err != nil {
		self.ServerAddr = old
		return fmt.Errorf("Could not set server address:\n%v", err)
	}

	return nil
}
