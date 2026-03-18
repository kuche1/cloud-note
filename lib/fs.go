package lib

import (
	"fmt"
	"os"
)

func WriteFileAtomic(file string, content []byte, temporaryFile string) error {
	err := os.WriteFile(temporaryFile, content, 0600)
	if err != nil {
		return fmt.Errorf("Could not create a new temporary file: %v", err)
	}

	err = os.Rename(temporaryFile, file)
	if err != nil {
		return fmt.Errorf("Could not replace old file with new one: %v", err)
	}

	return nil
}
