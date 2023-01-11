package main

import "store/filestore"

func main() {
	// new file store
	fs := &filestore.FileStore{"data"}

	// write to it
	id, err := fs.Set([]byte("some data"))
	_, _ = id, err

	// read an item
	data, err := fs.Get("xyz")
	_, _ = data, err

	// delete an item
	err = fs.Delete("xyz")
	_ = err
}
