package settings

import (
	"fmt"
)

func (self *Settings) SetServerAddr(new string) error {
	old := self.ServerAddr

	if old == new {
		// Be gentle with the SSD
		return nil
	}

	self.ServerAddr = new

	err := self.Save()
	if err != nil {
		self.ServerAddr = old
		return fmt.Errorf("Could not set server address:\n%v", err)
	}

	return nil
}

func (self *Settings) SetServerPassword(new string) error {
	old := self.ServerPassword

	if old == new {
		return nil
	}

	self.ServerPassword = new

	err := self.Save()
	if err != nil {
		self.ServerPassword = old
		return fmt.Errorf("Could not set server password:\n%v", err)
	}

	return nil
}

func (self *Settings) SetLastEditedNote(new string) error {
	old := self.ServerAddr

	if old == new {
		return nil
	}

	self.LastEditedNote = new

	// IMPROVE000: No need to rewrite all settings for a single item
	err := self.Save()
	if err != nil {
		self.LastEditedNote = old
		return fmt.Errorf("Could not set last edited server note:\n%v", err)
	}

	return nil
}
