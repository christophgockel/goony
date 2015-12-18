package files

import "os"

func OutputFile(name string) (File, error) {
	if name == "" {
		return os.Stdout, nil
	}

	return Filesystem.Create(name)
}
