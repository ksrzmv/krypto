package krypto

func Rotl(x, shift uint) uint {
	shift &= KR_WORD_SIZE - 1
	return (x << shift) | (x >> (KR_WORD_SIZE - shift))
}

func Rotr(x, shift uint) uint {
	shift &= KR_WORD_SIZE - 1
	return (x >> shift) | (x << (KR_WORD_SIZE - shift))
}

