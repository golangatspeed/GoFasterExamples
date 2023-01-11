package main

import (
	"store/store"
)

func main() {
	// new store, currently a file store
	st := store.NewStore("file")

	// write to it
	id, err := st.Write([]byte("some data"))
	_, _ = id, err

	// read an item
	data, err := st.Read("xyz")
	_, _ = data, err

	// delete an item
	err = st.Delete("xyz")
	_ = err
}