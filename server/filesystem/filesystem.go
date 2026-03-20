package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/kuche1/cloud-note/lib"
)

type Filesystem struct {
	storagePersistent string
	storageTemporary  string

	// IMPROVE000: Having a single lock for the whole filesystem sucks
	lock sync.RWMutex
}

func NewFilesystem(storageRoot string) (*Filesystem, error) {
	storagePersistent := filepath.Join(storageRoot, "persistent")
	err := os.MkdirAll(storagePersistent, 0755) // for some reason this does not work with 0600
	if err != nil {
		return nil, fmt.Errorf("Could not create persistent storage folder: %v", err)
	}

	storageTemporary := filepath.Join(storageRoot, "temporary")
	err = os.MkdirAll(storageTemporary, 0755) // for some reason this does not work with 0600
	if err != nil {
		return nil, fmt.Errorf("Could not create temporary storage folder: %v", err)
	}

	return &Filesystem{
		storagePersistent: storagePersistent,
		storageTemporary:  storageTemporary,
		lock:              sync.RWMutex{},
	}, nil
}

func (self *Filesystem) makePathLocal(path string) (_persistent string, _temporary string, _err error) {
	path = filepath.Clean(path)

	if !filepath.IsLocal(path) {
		return "", "", fmt.Errorf("Path is not local")
	}

	return filepath.Join(self.storagePersistent, path), filepath.Join(self.storageTemporary, path), nil
}

func (self *Filesystem) FileRead(unsafePath string) ([]byte, error) {
	persistent, _, err := self.makePathLocal(unsafePath)
	if err != nil {
		return nil, err
	}

	self.lock.RLock()
	defer self.lock.RUnlock()

	data, err := os.ReadFile(persistent)
	if err != nil {
		return nil, fmt.Errorf("Could not read file: %v", err)
	}

	return data, nil
}

func (self *Filesystem) FileWrite(unsafePath string, data []byte) error {
	persistent, temporary, err := self.makePathLocal(unsafePath)
	if err != nil {
		return err
	}

	self.lock.Lock()
	defer self.lock.Unlock()

	err = lib.WriteFileAtomic(persistent, data, temporary)
	if err != nil {
		return err
	}

	return nil
}

// IMPROVE000: ? Recursively collect all files
func (self *Filesystem) ListFiles() ([]string, error) {
	self.lock.RLock()

	entries, err := os.ReadDir(self.storagePersistent)
	if err != nil {
		self.lock.RUnlock()
		return nil, err
	}

	self.lock.RUnlock()

	ret := make([]string, 0, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ret = append(ret, entry.Name())
	}

	return ret, nil
}

func (self *Filesystem) FileCreateNew(unsafePath string) error {
	persistent, _, err := self.makePathLocal(unsafePath)
	if err != nil {
		return err
	}

	self.lock.Lock()
	defer self.lock.Unlock()

	file, err := os.OpenFile(persistent, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func (self *Filesystem) FileDeleteExisting(unsafePath string) error {
	persistent, _, err := self.makePathLocal(unsafePath)
	if err != nil {
		return err
	}

	self.lock.Lock()
	defer self.lock.Unlock()

	err = os.Remove(persistent)
	if err != nil {
		return err
	}

	return nil
}
