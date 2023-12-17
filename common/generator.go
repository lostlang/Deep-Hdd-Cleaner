package common

func LazyGenerateBytes(n int) *[]byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(i>>7) + byte(i<<3)
	}
	return &b
}
