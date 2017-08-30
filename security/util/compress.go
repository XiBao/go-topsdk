package util

func Compress(input []byte, toLen int) []byte {
	if toLen <= 0 {
		return nil
	}

	ret := make([]byte, toLen)
	for i, v := range input {
		ret[i%toLen] ^= v
	}
	return ret
}
