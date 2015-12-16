package files

import (
	"io"
	"os"
)

var Filesystem TheFileSystem = realFilesystem{}

type TheFileSystem interface {
	Open(name string) (File, error)
}

type File interface {
	io.Reader
}

type realFilesystem struct{}

func (realFilesystem) Open(name string) (File, error) {
	return os.Open(name)
}
