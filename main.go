package main

import (
	"bytes"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}

	var buffer = make([]byte, 8)
	var line string = ""

	for {

		_, err := file.Read(buffer)
		if err != nil {
			break
		}

		n := bytes.Index(buffer, []byte{10}) // 10 -> \n
		if n != -1 {

		}
	}
}
