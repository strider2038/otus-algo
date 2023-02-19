package main

import (
	"fmt"
	"io"
	"os"

	"github.com/strider2038/otus-algo/hw-20-run-length-encoding/rle"
)

const usageDescription = `RLE compressor.

Usage:
rle compress <input filename> <output filename> - to compress input filename and save it to output;
rle decompress <input filename> <output filename> - to decompress input filename and save it to output.
`

func main() {
	if len(os.Args) <= 1 || os.Args[1] == "help" {
		fmt.Println(usageDescription)

		return
	}

	switch os.Args[1] {
	case "compress", "decompress":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}
		if err := Execute(os.Args[1], os.Args[2], os.Args[3]); err != nil {
			fmt.Printf("%s error: %v\n", os.Args[1], err)
			os.Exit(1)
		}
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func Execute(command, inputFilename, outputFilename string) error {
	input, err := os.Open(inputFilename)
	if err != nil {
		return fmt.Errorf("open input: %w", err)
	}
	defer input.Close()

	output, err := os.Create(outputFilename)
	if err != nil {
		return fmt.Errorf("open output: %w", err)
	}
	defer output.Close()

	var f func(input io.Reader, output io.Writer) error

	switch command {
	case "compress":
		f = rle.Compress
	case "decompress":
		f = rle.Decompress
	default:
		return fmt.Errorf("unknown command")
	}

	if err := f(input, output); err != nil {
		return err
	}

	fmt.Printf("file '%s' %sed into '%s'", inputFilename, command, outputFilename)

	return nil
}
