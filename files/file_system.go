package files

import "os"

var Filesystem TheFileSystem = realFilesystem{}

type TheFileSystem interface {
	Create(name string) (File, error)
	Open(name string) (File, error)
}

type realFilesystem struct{}

func (realFilesystem) Open(name string) (File, error) {
	return os.Open(name)
}

func (realFilesystem) Create(name string) (File, error) {
	return os.Create(name)
}
