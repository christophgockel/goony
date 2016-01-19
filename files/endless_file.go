package files

import "io"

type EndlessFile struct {
	File
}

func (file EndlessFile) Read(b []byte) (int, error) {
	data, err := file.File.Read(b)

	if err == io.EOF {
		file.File.Seek(0, 0)
		data, err = file.File.Read(b)
	}

	return data, err
}
