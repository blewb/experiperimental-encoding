package codecs

import (
	"fmt"
)

type FiveBitCodec struct{}

func (FiveBitCodec) Encode(seq string) (string, error) {

	err := prevalidate(seq)
	if err != nil {
		return "", err
	}

	n := len(seq)

	bits := make([]byte, 0, (n/2)*5+1)
	bits = append(bits, '1')

	for s := 0; s < n-2; s += 2 {

		pattern := string(seq[s : s+2])

		if _, ok := fiveBitPatternToBinary[pattern]; !ok {
			return "", fmt.Errorf("invalid digit pattern in sequence - %s", pattern)
		}

		bits = append(bits, fiveBitPatternToBinary[pattern]...)

	}

	return binToHex(string(bits))

}

func (FiveBitCodec) Decode(seq string) (string, error) {
	return "", nil
}
