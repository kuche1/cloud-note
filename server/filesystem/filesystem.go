// TODO: Actually use

// TODO: Actually implement ability to interact with multiple files

package filesystem

import (
	"os"
	"sync"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/config"
)

type Filesystem struct {
	// storagePersistent string
	// storageTemporary  string
	// // TODO: Look into `sync.Map`

	rwLockForTheOnlySorryFileThatWeHave sync.RWMutex
}

func NewFilesystem(storageRoot string) *Filesystem {
	return &Filesystem{
		// storagePersistent:                   filepath.Join(storageRoot, "persistent"),
		// storageTemporary:                    filepath.Join(storageRoot, "temporary"),
		rwLockForTheOnlySorryFileThatWeHave: sync.RWMutex{},
	}
}

// func (self *Filesystem) makePathLocal(path string) (_persistent string, _temporary string, _err error) {
// 	path = filepath.Clean(path)

// 	if !filepath.IsLocal(path) {
// 		return "", fmt.Errorf("Path is not local")
// 	}

// 	return filepath.Join(self.storagePersistent, path), filepath.Join(self.storageTemporary, path), nil
// }

func (self *Filesystem) FileRead(unsafePath string) ([]byte, error) {
	// persistent, temporary, err := self.makePathLocal(unsafePath)
	// if err != nil {
	// 	return err
	// }

	self.rwLockForTheOnlySorryFileThatWeHave.RLock()
	defer self.rwLockForTheOnlySorryFileThatWeHave.RUnlock()

	data, err := os.ReadFile(config.NoteFile)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (self *Filesystem) FileWrite(unsafePath string, data []byte) error {
	// persistent, temporary, err := self.makePathLocal(unsafePath)
	// if err != nil {
	// 	return err
	// }

	self.rwLockForTheOnlySorryFileThatWeHave.RLock()
	defer self.rwLockForTheOnlySorryFileThatWeHave.RUnlock()

	err := lib.WriteFileAtomic(config.NoteFile, data, config.NoteFileTemporary)
	if err != nil {
		return err
	}

	return nil
}
