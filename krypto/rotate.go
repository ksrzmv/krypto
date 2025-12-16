package krypto

func Rotl(x, shift uint) uint {
	return ((x << (shift & KR_MODULUS)) | (x >> ((KR_WORD_SIZE - shift) & KR_MODULUS))) & KR_MODULUS
}

func Rotr(x, shift uint) uint {
	return ((x >> (shift & KR_MODULUS)) | (x << ((KR_WORD_SIZE - shift) & KR_MODULUS))) & KR_MODULUS
}

