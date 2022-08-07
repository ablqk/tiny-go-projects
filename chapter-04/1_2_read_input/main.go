package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/1_2_read_input/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	g := gordle.New(bufio.NewReader(os.Stdin))
	g.Play()
}
