package main

import (
	"fmt"
	"os"
)

type FooReader struct{}

func (f FooReader) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (f FooWriter) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}
func notMain() {
	var (
		reader FooReader
		writer FooWriter
	)
	input := make([]byte, 4096)
	s, _ := reader.Read(input)
	fmt.Printf("Read %d bytes from stdin\n", s)

	s, _ = writer.Write(input)
	fmt.Printf("Wrote %d bytes from stdin\n", s)

}
