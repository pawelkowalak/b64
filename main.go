package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide string to encode")
		os.Exit(1)
	}
	out := base64.StdEncoding.EncodeToString([]byte(os.Args[1]))
	fmt.Println(out)
}
