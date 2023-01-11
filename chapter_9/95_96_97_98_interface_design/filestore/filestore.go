package filestore

// FileStore is a struct representation of a file based store
type FileStore struct {
	Folder string
}

// Write writes a payload to the file store
func (fs *FileStore) Write(payload []byte) (string, error) {
	return "", nil
}

// Read retrieves an item of data from the store by id
func (fs *FileStore) Read(id string) ([]byte, error) {
	return []byte{}, nil
}

// Delete removes an item of data from the store by id
func (fs *FileStore) Delete(id string) error {
	return nil
}