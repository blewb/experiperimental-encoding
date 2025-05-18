package codecs

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
		pattern := int(seq[s]-'0')*5 + int(seq[s+1]-'0')
		bits = append(bits, numberToBinary(pattern, FIVE_BIT_WIDTH)...)
	}

	return binToHex(string(bits))

}

func (FiveBitCodec) Decode(seq string) (string, error) {
	return "", nil
}
