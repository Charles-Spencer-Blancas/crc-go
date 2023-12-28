package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var checkFlag = flag.Bool("c", false, "Check message validity")

func main() {
	usage := "Usage: [optional -c] [gen] [data]"
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("ERROR: Incorrect number of arguments\n", usage)
		os.Exit(1)
	}
	gen, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		fmt.Println("ERROR: Generator must be a number\n", usage)
	}

	data, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		fmt.Println("ERROR: Data must be a number\n", usage)
	}

	if *checkFlag {
		fmt.Println(checkCrc(gen, data))
	} else {
		fmt.Println(crc(gen, data))
	}
}
