package filestore

type FileLineStorage interface {
	Insert(line string) error
	Read() ([]string, error)
}
