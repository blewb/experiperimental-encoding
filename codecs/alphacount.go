package codecs

type AlphaCountCodec struct{}

func (AlphaCountCodec) Encode(seq string) (string, error) {

	err := prevalidate(seq)
	if err != nil {
		return "", err
	}

	n := len(seq)

	enc := make([]byte, 0, n)

	lastPattern := -1
	count := 0

	pushPattern := func(pt, ct int) {

		if ct > 10 {
			enc = append(enc, byte('0'+ct/10))
		}
		if ct%10 > 1 {
			enc = append(enc, byte('0'+ct%10))
		}
		enc = append(enc, byte('a'+pt))

	}

	for s := 0; s < n-1; s += 2 {

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

	return string(enc), nil

}

func (AlphaCountCodec) Decode(seq string) (string, error) {
	return "", nil
}
