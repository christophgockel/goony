package files

import "os"

func OpenForReading(name string) (File, error) {
	return Filesystem.Open(name)
}

func OpenForWriting(name string) (File, error) {
	if name == "" {
		return os.Stdout, nil
	}

	return Filesystem.Create(name)
}
