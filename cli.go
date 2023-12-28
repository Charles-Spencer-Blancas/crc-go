package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	checkFlag  = flag.Bool("c", false, "Check message validity")
	binaryMode = flag.Bool("b", false, "Parse input as binary and output binary")
	decodeFlag = flag.Bool("d", false, "Decode a message")
)

func main() {
	usage := "Usage: [opt -c xor -d] [opt -b] [gen] [data]"
	flag.Parse()
	if *checkFlag && *decodeFlag {
		fmt.Fprintln(os.Stderr, "ERROR: -c and -d flags are mutually exclusive")
		os.Exit(1)
	}
	base := 10
	if *binaryMode {
		base = 2
	}
	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "ERROR: Incorrect number of arguments\n", usage)
		os.Exit(1)
	}
	gen, err := strconv.ParseUint(args[0], base, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Generator must be a number >= 2\n", usage)
	}

	data, err := strconv.ParseUint(args[1], base, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Data must be a number >= 0\n", usage)
	}

	if *checkFlag {
		out, err := checkCrc(gen, data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(out)
	}

	out := "%d\n"
	if *binaryMode {
		out = "%b\n"
	}

	if *decodeFlag {
		msg, err := decodeCrc(gen, data)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf(out, msg)
	} else {
		res, err := crc(gen, data)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf(out, res)
	}
}
