package files

import "io"

type EndlessFile struct {
	File
}

func (file EndlessFile) Read(b []byte) (int, error) {
	data, e := file.File.Read(b)

	if e == io.EOF {
		file.File.Seek(0, 0)
		data, e = file.File.Read(b)
	}

	return data, e
}
