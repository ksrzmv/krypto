// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package krypto

func Encrypt(data []byte, key []byte) []byte {
	prepData := dataToUintArray(data, Enc)
	S := keyExpansion(key)

	for i := 0; i < len(prepData) - 1; i += 2 {
		A := prepData[i]
		B := prepData[i+1]
		A += S[0]
		B += S[1]
		for j := 1; j <= KR_ROUNDS; j++ {
			A = A ^ B
			A = Rotl(A, B)
			A += S[2*j]
			B = B ^ A
			B = Rotl(B, A)
			B += S[2*j+1]
		}
		prepData[i] = A
		prepData[i+1] = B
	}

	return dataFromUintArray(prepData, Enc)
}

func Decrypt(data []byte, key []byte) []byte {
	prepData := dataToUintArray(data, Dec)
	S := keyExpansion(key)

	for i := 0; i < len(prepData) - 1; i += 2 {
		A := prepData[i]
		B := prepData[i+1]
		for j := KR_ROUNDS; j >= 1; j-- {
			B -= S[2*j+1]
			B = Rotr(B, A)
			B ^= A
			A -= S[2*j]
			A = Rotr(A, B)
			A ^= B
		}
		B -= S[1]
		A -= S[0]
		prepData[i+1] = B
		prepData[i] = A
	}

	return dataFromUintArray(prepData, Dec)
}

