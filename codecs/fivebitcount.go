package codecs

import (
	"fmt"
)

type FiveBitCountCodec struct{}

func (FiveBitCountCodec) Encode(seq string) (string, error) {

	err := prevalidate(seq)
	if err != nil {
		return "", err
	}

	n := len(seq)

	// 3 bit count + 5 bit pattern for half the sequence
	// (3 + 5) / 2 = 4
	bits := make([]byte, 0, n*4+1)
	bits = append(bits, '1')

	lastPattern := ""
	count := 0

	pushPattern := func(pt string, ct int) {

		for ct > 7 {
			bits = append(bits, threeBitDigitToBinary[7]...)
			bits = append(bits, fiveBitPatternToBinary[pt]...)
			ct -= 7
		}

		bits = append(bits, threeBitDigitToBinary[ct]...)
		bits = append(bits, fiveBitPatternToBinary[pt]...)

	}

	for s := 0; s < n-2; s += 2 {

		pattern := string(seq[s : s+2])

		if _, ok := fiveBitPatternToBinary[pattern]; !ok {
			return "", fmt.Errorf("invalid digit pattern in sequence - %s", pattern)
		}

		if pattern == lastPattern {
			count++
			continue
		}

		if lastPattern != "" {
			pushPattern(lastPattern, count)
		}

		lastPattern = pattern
		count = 1

	}

	pushPattern(lastPattern, count)

	return binToHex(string(bits))

}

func (FiveBitCountCodec) Decode(seq string) (string, error) {
	return "", nil
}
