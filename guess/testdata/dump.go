package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	sep := ""
	buf := bufio.NewReader(file)
	for {
		b, err := buf.ReadByte()
		if err != nil {
			break
		}
		fmt.Printf("%s0x%02x", sep, b)
		sep = ", "
	}
	fmt.Println()
}
