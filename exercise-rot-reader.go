package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	len, err := r13.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = (c-'a'+13)%26 + 'a'
		}
		if c >= 'A' && c <= 'Z' {
			b[i] = (c-'A'+13)%26 + 'A'
		}
	}
	return len, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
