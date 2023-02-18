package main

import (
	"errors"
	"testing"
)

var ERR_WOULD_OVERFLOW = errors.New("would overflow")

func Adder(list ...uint8) (uint8, error) {
	var t uint8
	for i := 0; i < len(list); i++ {
		if int(t)+int(list[i]) > 255 {
			return 0, ERR_WOULD_OVERFLOW
		}
		t += list[i]
	}

	return t, nil
}

func TestAdder(t *testing.T) {
	testCases := []struct {
		name    string
		inputs  []uint8
		wantRes uint8
		wantErr error
	}{
		{
			"simple total",
			[]uint8{1, 1, 7},
			9,
			nil,
		},
		{
			"overflow prevented",
			[]uint8{255, 1},
			0,
			ERR_WOULD_OVERFLOW,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRes, gotErr := Adder(tc.inputs...)
			if gotRes != tc.wantRes {
				t.Errorf("wanted %v got %v", tc.wantRes, gotRes)
			}
			if gotErr != tc.wantErr {
				t.Errorf("wanted error %v got %v", tc.wantErr, gotErr)
			}
		})
	}
}
