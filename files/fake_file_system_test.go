package files_test

import (
	"bytes"
	"github.com/christophgockel/goony/files"
	"os"
)

var fakeFilesystem aFakeFileSystem = aFakeFileSystem{}

type aFakeFileSystem struct {
	files.TheFileSystem

	fileExists                bool
	fileHasCorrectPermissions bool

	openHasBeenCalled   bool
	createHasBeenCalled bool
}

func (fs *aFakeFileSystem) Open(name string) (files.File, error) {
	fs.openHasBeenCalled = true

	return fs.fakeFileBehavior(name)
}

func (fs *aFakeFileSystem) Create(name string) (files.File, error) {
	fs.createHasBeenCalled = true

	return fs.fakeFileBehavior(name)
}

func (fs aFakeFileSystem) fakeFileBehavior(name string) (files.File, error) {
	if fs.fileExists == false {
		return nil, os.ErrNotExist
	}

	if fs.fileHasCorrectPermissions == false {
		return nil, os.ErrPermission
	}

	return fakeFile{}, nil
}

func prepareFilesystemWithAccessibleFile() {
	files.Filesystem = &fakeFilesystem

	fakeFilesystem.fileExists = true
	fakeFilesystem.fileHasCorrectPermissions = true
}

func prepareFilesystemWithUnexistingFile() {
	files.Filesystem = &fakeFilesystem

	fakeFilesystem.fileExists = false
	fakeFilesystem.fileHasCorrectPermissions = true
}

func prepareFilesystemWithUnaccessibleFile() {
	files.Filesystem = &fakeFilesystem

	fakeFilesystem.fileExists = true
	fakeFilesystem.fileHasCorrectPermissions = false
}

func fileIsWritable() bool {
	return fakeFilesystem.createHasBeenCalled
}

func fileIsReadable() bool {
	return fakeFilesystem.openHasBeenCalled
}

type fakeFile struct {
	files.File

	lines   []string
	content *bytes.Buffer
}

func (file fakeFile) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (file fakeFile) Read(p []byte) (n int, err error) {
	return file.content.Read(p)
}

func (file fakeFile) Seek(offset int64, whence int) (int64, error) {
	for _, line := range file.lines {
		file.AddLine(line)
	}

	return 0, nil
}

func (file *fakeFile) AddLine(line string) {
	if file.content == nil {
		file.content = new(bytes.Buffer)
	}
	file.lines = append(file.lines, line)
	file.content.WriteString(line + "\n")
}
