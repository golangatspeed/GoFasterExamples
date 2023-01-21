package main

import (
	"errors"
	"fmt"
	"reflect"
)

func sumIntAny(x, y interface{}) (interface{}, error) {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		return nil, errors.New("mismatched  types")
	}

	switch v := x.(type) {
	case int8:
		return v + y.(int8), nil
	case int16:
		return v + y.(int16), nil
	case int32:
		return v + y.(int32), nil
	case int64:
		return v + y.(int64), nil
	}

	return nil, nil
}

func main() {
	res, err := sumIntAny(int8(1), int8(2))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result:", res)
	//fmt.Println("Addition:", res+res)
}