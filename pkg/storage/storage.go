package storage

type Storage interface {
	Write(data []string) error
	Read() ([]byte, error)
}
