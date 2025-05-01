package main

import (
	"fmt"

	"github.com/blewb/experiperimental-encoding/codecs"
)

const SEQ = "0004000000440000000030000011310002111120121122111211121212121222121111210111111000111100"

const TEST_SEQ = "01010101"

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
	}

	fmt.Println("")

	for i, item := range codecList {

		fmt.Printf("#%d %s\n", i+1, item.name)

		encoded, err := item.codec.Encode(SEQ)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		fmt.Printf("%s (%d)\n", encoded, len(encoded))

		fmt.Println("")

	}

}
