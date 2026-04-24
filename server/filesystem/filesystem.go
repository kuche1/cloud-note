package filesystem

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/kuche1/cloud-note/lib"
)

const _IncorrectPasswordMessage string = "Incorrect password"

type Filesystem struct {
	storagePersistent string
	storageTemporary  string

	// IMPROVE000: Having a single lock for the whole filesystem sucks
	lock sync.RWMutex

	passwordFile          string
	passwordFileTemporary string
	passwordLock          sync.RWMutex
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
		storagePersistent:     storagePersistent,
		storageTemporary:      storageTemporary,
		lock:                  sync.RWMutex{},
		passwordFile:          filepath.Join(storageRoot, "password"),
		passwordFileTemporary: filepath.Join(storageRoot, "password-tmp"),
		passwordLock:          sync.RWMutex{},
	}, nil
}

func (self *Filesystem) makePathLocal(path string) (_persistent string, _temporary string, _err error) {
	path = filepath.Clean(path)

	if !filepath.IsLocal(path) {
		return "", "", fmt.Errorf("Path is not local")
	}

	return filepath.Join(self.storagePersistent, path), filepath.Join(self.storageTemporary, path), nil
}

func (self *Filesystem) readPasswordFileUnsafe() (string, error) {
	passwordBytes, err := os.ReadFile(self.passwordFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", fmt.Errorf("Could not read password file: %v", err)
	}

	return string(passwordBytes), nil
}

func (self *Filesystem) CheckPassword(userPassword string) error {
	if userPassword == "" {
		return fmt.Errorf("Empty passwords are not accepted")
	}

	self.passwordLock.RLock()

	actualPassword, err := self.readPasswordFileUnsafe()
	if err != nil {
		self.passwordLock.RUnlock()
		return err
	}

	if actualPassword == "" {
		self.passwordLock.RUnlock()

		self.passwordLock.Lock()
		defer self.passwordLock.Unlock()

		doubleCheckPassword, err := self.readPasswordFileUnsafe()
		if doubleCheckPassword != actualPassword {
			// Another thread must have changed the password while we were waiting
			// for the lock
			return fmt.Errorf(_IncorrectPasswordMessage)
		}

		err = lib.FileWriteAtomic(
			self.passwordFile,
			[]byte(userPassword),
			self.passwordFileTemporary,
		)
		if err != nil {
			return fmt.Errorf("Could not save new password file: %v", err)
		}

		return nil
	}

	defer self.passwordLock.RUnlock()

	if actualPassword != userPassword {
		return fmt.Errorf(_IncorrectPasswordMessage)
	}

	return nil
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

	err = lib.FileWriteAtomic(persistent, data, temporary)
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

func (self *Filesystem) FileRename(unsafePathOld string, unsafePathNew string) (_refusal string, _err error) {
	pathOld, _, err := self.makePathLocal(unsafePathOld)
	if err != nil {
		return "", err
	}

	pathNew, _, err := self.makePathLocal(unsafePathNew)
	if err != nil {
		return "", err
	}

	self.lock.Lock()
	defer self.lock.Unlock()

	_, err = os.Stat(pathNew)
	if err == nil {
		return "Destination already exists", nil
	}

	err = os.Rename(pathOld, pathNew)
	if err != nil {
		return "", err
	}

	return "", nil
}
