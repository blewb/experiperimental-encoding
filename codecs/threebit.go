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

		bits = append(bits, threeBitDigitToBinary[digit]...)

	}

	return binToHex(string(bits))

}

func (ThreeBitCodec) Decode(seq string) (string, error) {
	return "", nil
}
