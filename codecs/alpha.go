package codecs

type AlphaCodec struct{}

func (AlphaCodec) Encode(seq string) (string, error) {

	err := prevalidate(seq)
	if err != nil {
		return "", err
	}

	n := len(seq)

	enc := make([]byte, 0, n/2)

	for s := 0; s < n-1; s += 2 {
		pattern := int(seq[s]-'0')*5 + int(seq[s+1]-'0')
		enc = append(enc, byte('a'+pattern))
	}

	return string(enc), nil

}

func (AlphaCodec) Decode(seq string) (string, error) {

	return "", nil

}
