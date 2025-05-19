package codecs

import "fmt"

type FiveBitCodec struct{}

func (FiveBitCodec) Encode(seq string) (string, error) {

	err := prevalidate(seq)
	if err != nil {
		return "", err
	}

	n := len(seq)

	bits := make([]byte, 0, (n/2)*5+1)
	bits = append(bits, '1')

	for s := 0; s < n-1; s += 2 {
		pattern := int(seq[s]-'0')*5 + int(seq[s+1]-'0')
		bits = append(bits, numberToBinary(pattern, FIVE_BIT_WIDTH)...)
	}

	return binToHex(string(bits))

}

func (FiveBitCodec) Decode(seq string) (string, error) {

	bin, _ := hexToBin(seq)
	output := make([]byte, 0, SEQUENCE_LENGTH)

	// Start at 1 to ignore the leading '1' bit
	for s := 1; s < len(bin); s += FIVE_BIT_WIDTH {

		unit := string(bin[s : s+FIVE_BIT_WIDTH])
		pattern := binaryToNumber(unit)

		if pattern < 0 || pattern >= 25 {
			return "", fmt.Errorf("invalid pattern: %b", pattern)
		}

		output = append(output, byte('0'+pattern/5))
		output = append(output, byte('0'+pattern%5))

	}

	return string(output), nil

}
