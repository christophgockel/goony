package files

func Open(name string) (File, error) {
	return Filesystem.Open(name)
}
