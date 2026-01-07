// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package krypto

func Encrypt(data []byte, key []byte) []byte {
	prep_data := dataToUintArray(data, Enc)
	S := keyExpansion(key)

	// enrich data with IV
	iv_A := GenerateKey(KR_WORD_SIZE_BYTES)
	iv_B := GenerateKey(KR_WORD_SIZE_BYTES)

	iv_data_length := len(prep_data) + 2
	iv_data := make([]uint, iv_data_length)
	for i := 0; i < len(iv_A); i++ {
		iv_data[0] += uint(iv_A[i] << (56 - i * KR_WORD_SIZE_BYTES))
		iv_data[1] += uint(iv_B[i] << (56 - i * KR_WORD_SIZE_BYTES))
	}
	for idx, val := range prep_data {
		iv_data[idx+2] = val
	}
	prep_data = iv_data
	// -----


	for i := 0; i < len(prep_data) - 1; i += 2 {
		A := prep_data[i]
		B := prep_data[i+1]

		if i > 1 {
			A ^= prep_data[i-2]
			B ^= prep_data[i-1]
		}

		// start block encryption
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
		// end block encryption

		prep_data[i] = A
		prep_data[i+1] = B
	}

	return dataFromUintArray(prep_data, Enc)
}

func Decrypt(data []byte, key []byte) []byte {
	prep_data := dataToUintArray(data, Dec)
	S := keyExpansion(key)

	feedback_A := uint(0)
	feedback_B := uint(0)
	var prev_A, prev_B uint

	for i := 0; i < len(prep_data) - 1; i += 2 {
		A := prep_data[i]
		B := prep_data[i+1]

		prev_A = A
		prev_B = B

		// start block decryption
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
		// end block decryption

	 	B ^= feedback_B
	 	A ^= feedback_A

		feedback_A = prev_A
		feedback_B = prev_B

		prep_data[i+1] = B
		prep_data[i] = A
	}

	return dataFromUintArray(prep_data[2:], Dec)
}

