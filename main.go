package main

import "github.com/blewb/experiperimental-encoding/codecs"

const SEQ = "0004000000440000000030000011310002111120121122111211121212121222121111210111111000111100"

const TEST_SEQ = "01234"

func main() {

	codec := codecs.ThreeBitCountCodec{}

	codec.Encode(SEQ)

}
