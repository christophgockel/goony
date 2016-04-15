package request

type Data struct {
	Host     string
	Path     string
	Username string
	Password string
}

func (data Data) Url() string {
	return data.Host + data.Path
}
