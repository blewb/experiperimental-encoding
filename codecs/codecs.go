package codecs

import (
	"fmt"
	"math/big"
	"strconv"
)

// Magic number alert
const SEQUENCE_LENGTH = 88

// Best practices and all that...
const THREE_BIT_WIDTH = 3
const FIVE_BIT_WIDTH = 5

type Codec interface {
	Encode(string) (string, error)
	Decode(string) (string, error)
}

func prevalidate(seq string) error {

	// if len(seq) != SEQUENCE_LENGTH {
	// 	return fmt.Errorf("invalid sequence length %d - require %d - %s", len(seq), SEQUENCE_LENGTH, seq)
	// }

	return nil

}

func validDigit(x int) bool {

	if x < 0 || x > 4 {
		return false
	}

	return true

}

func binToHex(bits string) (string, error) {

	number := big.NewInt(0)
	number.SetString(bits, 2)

	return fmt.Sprintf("%x", number), nil

}

func hexToBin(hex string) (string, error) {

	number := big.NewInt(0)
	number.SetString(hex, 16)

	return fmt.Sprintf("%b", number), nil

}

func numberToBinary(num, bits int) string {

	if bits < 1 || bits > 8 {
		return ""
	}

	format := fmt.Sprintf("%%0%db", bits)

	return fmt.Sprintf(format, num)

}

func binaryToNumber(bits string) int {

	num, err := strconv.ParseInt(bits, 2, 32)
	if err != nil {
		return -1
	}

	return int(num)

}
