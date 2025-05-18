package codecs

import (
	"fmt"
)

type ThreeBitCodec struct{}

func (ThreeBitCodec) Encode(seq string) (string, error) {

	err := prevalidate(seq)
	if err != nil {
		return "", err
	}

	bits := make([]byte, 0, SEQUENCE_LENGTH*3+1)
	bits = append(bits, '1')

	for _, s := range seq {

		digit := int(s - '0')

		if !validDigit(digit) {
			return "", fmt.Errorf("invalid digit in sequence - %d", digit)
		}

		bits = append(bits, numberToBinary(digit, THREE_BIT_WIDTH)...)

	}

	return binToHex(string(bits))

}

func (ThreeBitCodec) Decode(seq string) (string, error) {

	bin, _ := hexToBin(seq)
	output := make([]byte, 0, SEQUENCE_LENGTH)

	// Start at 1 to ignore the leading '1' bit
	for s := 1; s < len(bin); s += THREE_BIT_WIDTH {

		unit := string(bin[s : s+THREE_BIT_WIDTH])
		digit := binaryToNumber(unit)

		if !validDigit(digit) {
			return "", fmt.Errorf("invalid digit in sequence - %d", digit)
		}

		output = append(output, byte('0'+digit))

	}

	return string(output), nil

}
