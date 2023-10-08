package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	default:
		return r
	}
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	if b == nil || len(b) == 0 {
		return 0, nil
	}
	read, err := rot.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < read; i++ {
		b[i] = byte(rot13(rune(b[i])))
	}
	return read, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	ori := strings.NewReader("123456")
	b := make([]byte, 5)
	read, err := ori.Read(b)
	fmt.Printf("%d, %v, %s\n", read, err, b)
	fmt.Println("-----------------------------")
	read, err = ori.Read(b)
	fmt.Printf("%d, %v, %s\n", read, err, b)
	fmt.Println("-----------------------------")
	read, err = ori.Read(b)
	fmt.Printf("%d, %v, %s\n", read, err, b)

}
