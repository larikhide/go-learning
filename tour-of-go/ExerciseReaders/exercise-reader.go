package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

/* func (mr MyReader) Read(b []byte) (int, error) {
	for k := range b {
		b[k] = 'A'
	}
	return 1, nil
} */

func (mr MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
