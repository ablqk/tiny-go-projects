package main

import (
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/1_2_read_input/gordle"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}
