package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {

	var chn = make(chan string)

	go func() {
		var buf = make([]byte, 8)
		var line string = ""

		for {
			n, err := f.Read(buf)
			if err != nil && err != io.EOF {
				break
			}

			if n == 0 {
				break
			}

			chunk := buf[:n]

			for {
				sep := bytes.IndexByte(chunk, '\n')

				if sep == -1 {
					line += string(chunk)
					break
				} else {
					line += string(chunk[:sep])
					chn <- line
					line = ""
					chunk = chunk[sep+1:]
				}
			}
			if err == io.EOF {
				break
			}
		}
		if line != "" {
			chn <- line
		}
		close(chn)
	}()
	return chn
}

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for line := range getLinesChannel(file) {
		fmt.Printf("read: %s\n", line)
	}
}
