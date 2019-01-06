package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

func main() {
	text, decode, err := readArgs(os.Args)
	if err != nil {
		fmt.Println("Usage: b64 [-d] <value>")
		os.Exit(1)
	}

	var out string
	if decode {
		bb, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			fmt.Printf("Error during base64 decode: %v", err)
			os.Exit(1)
		}
		out = string(bb)
	} else {
		out = base64.StdEncoding.EncodeToString([]byte(text))
	}
	fmt.Println(out)
}

func readArgs(args []string) (string, bool, error) {
	switch len(args) {
	case 2:
		return args[1], false, nil
	case 3:
		if args[1] == "-d" {
			return args[2], true, nil
		}
	}
	return "", false, errors.New("invalid input arguments")
}
