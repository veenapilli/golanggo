package main

import (
	"fmt"
	"io"
	"net/http"
)

type logger struct{}

func main() {
	resp, _ := http.Get("http://google.com")
	fmt.Println(resp.StatusCode)

	/*
		// creating byte slice with known capacity
		bs := make([]byte, 99999)
		// read func can not extend the slice. it will only buffer till the capacity provided initially

		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/
	// io.Copy(os.Stdout, resp.Body)
	/**
	stdout - console
	copy date from reader to the writer interface.
	os.stdout for eg implements the writer interface.
	resp.Body gives the reader interface for the writer
	*/

	lo := logger{}
	count, _ := io.Copy(lo, resp.Body)
	fmt.Println("Bytes: " + string(count))
}

/**
implementing the writer interface
*/
func (logger) Write(bs []byte) (int, error) {
	// fmt.Println(string(bs))
	return len(bs), nil
}
