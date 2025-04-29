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
			bits = append(bits, threeBitDigitToBinary[7]...)
			bits = append(bits, threeBitDigitToBinary[dg]...)
			ct -= 7
		}

		bits = append(bits, threeBitDigitToBinary[ct]...)
		bits = append(bits, threeBitDigitToBinary[dg]...)

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
