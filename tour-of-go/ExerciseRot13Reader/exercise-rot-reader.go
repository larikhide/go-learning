package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for k := range p {
		p[k] = rot13(p[k])
	}
	return
}

func rot13(b byte) byte {
	switch {
	case b >= 'A' && b <= 'Z':
		return (b-'A'+13)%('Z'-'A'+1) + 'A'
	case b >= 'a' && b <= 'z':
		return (b-'a'+13)%('z'-'a'+1) + 'a'
	default:
		return b //not a letter
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
