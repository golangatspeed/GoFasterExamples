package store

import "store/filestore"

// ReadWriteDeleter must be implemented for each storage mechanism
type ReadWriteDeleter interface {
	Read(id string) (payload []byte, err error)
	Write(payload []byte) (id string, err error)
	Delete(id string) error
}

// NewStore creates a store from available stores, depending on storeType argument
func NewStore(storeType string) ReadWriteDeleter {
	var s ReadWriteDeleter
	switch storeType {
	case "file":
		s = &filestore.FileStore{
			Folder: "data",
		}
	}
	return s
}
