package main

import (
	"fmt"
	"reflect"
)

func MakeAddToMap[K comparable, V any](k K, v V) map[K]V {
	mp := make(map[K]V)
	mp[k] = v
	return mp
}

func main() {
	mp := MakeAddToMap("key", 1)
	fmt.Println(reflect.TypeOf(mp))
	fmt.Println(mp)
}