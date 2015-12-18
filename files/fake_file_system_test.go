package files_test

import (
	"github.com/christophgockel/goony/files"
	"os"
)

var fakeFilesystem aFakeFileSystem = aFakeFileSystem{}

type aFakeFileSystem struct {
	files.TheFileSystem

	fileExists                bool
	fileHasCorrectPermissions bool
}

type fakeFile struct {
	files.File
}

func (fs aFakeFileSystem) Open(name string) (files.File, error) {
	return fs.fakeFileBehavior(name)
}

func (fs aFakeFileSystem) Create(name string) (files.File, error) {
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
