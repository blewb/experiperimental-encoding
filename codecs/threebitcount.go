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
			fmt.Println(count, "x", lastDigit)
		}

		lastDigit = digit
		count = 1

		// bits = append(bits, fmt.Sprintf("%03b", digit)...)

	}

	fmt.Println(count, " x ", lastDigit)

	return binToHex(string(bits))

}

func (ThreeBitCountCodec) Decode(seq string) (string, error) {
	return "", nil
}
