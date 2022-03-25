package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	io.Pipe()

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	pipedInText := strings.TrimSpace(string(bytes))
	fmt.Println(pipedInText)
}
