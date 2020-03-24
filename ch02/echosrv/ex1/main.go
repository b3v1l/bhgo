package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

type FooWriter struct{}

func (f *FooReader) Read(b []byte) (int, error) {

	fmt.Print("IN> ")
	return os.Stdin.Read(b)
}

func (f *FooWriter) Write(b []byte) (int, error) {

	fmt.Print("OUT> ")
	return os.Stdout.Write(b)
}

func main() {

	var (
		reader FooReader
		writer FooWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Error in input/output data")

	}
}

//	input := make([]byte, 4096)
//
//	r, err := reader.Read(input)
//	if err != nil {
//		log.Fatalln("Error with input data")
//	}
//	fmt.Println(r)
//
//	w, err := writer.Write(input)
//	if err != nil {
//		log.Fatalln("Error with output data")
//	}
//
//	fmt.Println(w)
