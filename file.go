package main

import (
	"errors"
	"os"
)

type File struct {
	fileName string
	file     *os.File
}

func CreateFile(fileName string) (*File, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return &File{
		fileName: fileName,
		file:     f,
	}, nil
}

func (f *File) Defer() error {
	if f == nil {
		return nil
	}

	return f.file.Close()
}

func (f *File) Sync() error {
	if f == nil {
		return errors.New("File is not initialized")
	}

	return f.file.Sync()
}

func (f *File) Write(s string) error {
	if f == nil {
		return errors.New("File is not initialized")
	}

	_, err := f.file.WriteString(s)
	return err
}

func (f *File) WriteSlice(sl []string) error {
	for _, s := range sl {
		if err := f.Write(s); err != nil {
			return err
		}
	}

	return f.Sync()
}
