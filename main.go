package main

import (
	"fmt"

	"github.com/blewb/experiperimental-encoding/codecs"
)

const SEQ = "0004000000440000000030000011310002111120121122111211121212121222121111210111111000111100"

const TEST_SEQ = "01234444"

func main() {

	codecList := []struct {
		codec codecs.Codec
		name  string
	}{
		{
			codecs.ThreeBitCodec{},
			"Three Bit",
		},
		{
			codecs.ThreeBitCountCodec{},
			"Three Bit Count",
		},
		{
			codecs.FiveBitCodec{},
			"Five Bit",
		},
		{
			codecs.FiveBitCountCodec{},
			"Five Bit Count",
		},
		{
			codecs.AlphaCodec{},
			"Alpha",
		},
		{
			codecs.AlphaCountCodec{},
			"Alpha Count",
		},
	}

	sequence := SEQ

	fmt.Println("")

	fmt.Println("####........")
	fmt.Println("###..", sequence)
	fmt.Println("##........")

	fmt.Println("")

	for i, item := range codecList {

		// Name
		fmt.Printf("#%d %s\n", i+1, item.name)

		encoded, err := item.codec.Encode(sequence)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		// Encoded
		fmt.Printf("%s (%d)\n", encoded, len(encoded))

		decoded, err := item.codec.Decode(encoded)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		// Decoder Issue
		if decoded != sequence {
			fmt.Printf("decoder does not match input: %s\n", decoded)
		}

		fmt.Println("")

	}

}
