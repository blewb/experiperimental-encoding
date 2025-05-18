package codecs

import (
	"fmt"
)

type ThreeBitCountCodec struct{}

func (ThreeBitCountCodec) Encode(seq string) (string, error) {

	err := prevalidate(seq)
	if err != nil {
		return "", err
	}

	// Worst case scenario: no repeated digits
	// E.g. 01234 becomes 1011121314
	bits := make([]byte, 0, SEQUENCE_LENGTH*6+1)
	bits = append(bits, '1')

	lastDigit, count := -1, 0

	pushDigits := func(dg, ct int) {

		for ct > 7 {
			bits = append(bits, numberToBinary(7, THREE_BIT_WIDTH)...)
			bits = append(bits, numberToBinary(dg, THREE_BIT_WIDTH)...)
			ct -= 7
		}

		bits = append(bits, numberToBinary(ct, THREE_BIT_WIDTH)...)
		bits = append(bits, numberToBinary(dg, THREE_BIT_WIDTH)...)

	}

	for _, s := range seq {

		digit := int(s - '0')

		if !validDigit(digit) {
			return "", fmt.Errorf("invalid digit in sequence - %d", digit)
		}

		if digit == lastDigit {
			count++
			continue
		}

		if lastDigit >= 0 {
			pushDigits(lastDigit, count)
		}

		lastDigit = digit
		count = 1

	}

	pushDigits(lastDigit, count)

	return binToHex(string(bits))

}

func (ThreeBitCountCodec) Decode(seq string) (string, error) {
	return "", nil
}
