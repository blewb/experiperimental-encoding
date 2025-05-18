package codecs

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

	lastPattern := -1
	count := 0

	pushPattern := func(pt, ct int) {

		for ct > 7 {
			bits = append(bits, numberToBinary(7, THREE_BIT_WIDTH)...)
			bits = append(bits, numberToBinary(pt, FIVE_BIT_WIDTH)...)
			ct -= 7
		}

		bits = append(bits, numberToBinary(ct, THREE_BIT_WIDTH)...)
		bits = append(bits, numberToBinary(pt, FIVE_BIT_WIDTH)...)

	}

	for s := 0; s < n-2; s += 2 {

		pattern := int(seq[s]-'0')*5 + int(seq[s+1]-'0')

		if pattern == lastPattern {
			count++
			continue
		}

		if lastPattern >= 0 {
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
