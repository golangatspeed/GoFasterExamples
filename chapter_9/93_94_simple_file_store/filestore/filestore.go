package filestore

// FileStore is a struct representation of a file based store
type FileStore struct {
	Folder string
}

// Set writes a payload to the file store
func (fs *FileStore) Set(payload []byte) (string, error) {
	return "", nil
}

// Get retrieves an item of data from the store by id
func (fs *FileStore) Get(id string) ([]byte, error) {
	return []byte{}, nil
}

// Delete removes an item of data from the store by id
func (fs *FileStore) Delete(id string) error {
	return nil
}