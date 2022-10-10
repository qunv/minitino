package helpers

import (
	"bytes"
	"os"
)

func ReadFile(filePath string) (*bytes.Buffer, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(file)
	return buff, nil
}

func WriteFile(destination string, b *bytes.Buffer) error {
	return os.WriteFile(destination, b.Bytes(), 0644)
}

func ReadDir(dir string) ([]os.DirEntry, error) {
	return os.ReadDir(dir)
}
