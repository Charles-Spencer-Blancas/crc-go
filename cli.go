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
)

func main() {
	usage := "Usage: [optional -c] [gen] [data]"
	flag.Parse()
	base := 10
	if *binaryMode {
		base = 2
	}
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("ERROR: Incorrect number of arguments\n", usage)
		os.Exit(1)
	}
	gen, err := strconv.ParseUint(args[0], base, 64)
	if err != nil {
		fmt.Println("ERROR: Generator must be a number >= 2\n", usage)
	}

	data, err := strconv.ParseUint(args[1], base, 64)
	if err != nil {
		fmt.Println("ERROR: Data must be a number >= 0\n", usage)
	}

	if *checkFlag {
		out, err := checkCrc(gen, data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(out)
	} else {
		out := "%d\n"
		if *binaryMode {
			out = "%b\n"
		}
		res, err := crc(gen, data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf(out, res)
	}
}
