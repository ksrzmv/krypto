// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package krypto

func Encrypt(data []byte, key []byte) ([]byte, error) {
	// convert byte array to uint blocks
	prep_data, err := dataToUintArray(data, Enc)
	if err != nil {
		return nil, err
	}

	// create expansion key array
	S := keyExpansion(key)

	// enrich data with initialization vector (IV)

	// generate dword of random data for first IV block
	iv, err := GenerateKey(KR_DWORD_SIZE_BYTES)
	if err != nil {
		return nil, err
	}

	// length of enriched data = length of actual data + 2 iv blocks
	iv_data_length := len(prep_data) + 2
	iv_data := make([]uint, iv_data_length)
	for i := 0; i < len(iv)-KR_WORD_SIZE_BYTES; i++ {
		iv_data[0] += uint(iv[i] << (56 - i * KR_WORD_SIZE_BYTES))
		iv_data[1] += uint(iv[i+KR_WORD_SIZE_BYTES] << (56 - i * KR_WORD_SIZE_BYTES))
	}
	for idx, val := range prep_data {
		iv_data[idx+2] = val
	}
	prep_data = iv_data
	// -----


	// range over all two-block batches
	for i := 0; i < len(prep_data) - 1; i += 2 {
		A := prep_data[i]
		B := prep_data[i+1]

		if i > 1 {
			A ^= prep_data[i-2]
			B ^= prep_data[i-1]
		}

		// rc5 block encryption
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

	// convert uint blocks to byte array
	result, err := dataFromUintArray(prep_data, Enc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Decrypt(data []byte, key []byte) ([]byte, error) {
	// convert byte array to uint blocks
	prep_data, err := dataToUintArray(data, Dec)
	if err != nil {
		return nil, err
	}

	// create expansion key array
	S := keyExpansion(key)

	// CBC mode feedback; initialized as 0 for first XOR
	feedback_A := uint(0)
	feedback_B := uint(0)
	var prev_A, prev_B uint

	// range over all two-block batches
	for i := 0; i < len(prep_data) - 1; i += 2 {
		A := prep_data[i]
		B := prep_data[i+1]

		prev_A = A
		prev_B = B

		// rc5 block decryption
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

		// feedback for next round is previous encrypted two-blocks batch
		feedback_A = prev_A
		feedback_B = prev_B

		prep_data[i+1] = B
		prep_data[i] = A
	}

	// convert uint blocks to byte array
	result, err := dataFromUintArray(prep_data[2:], Dec)
	if err != nil {
		return nil, err
	}

	return result, nil
}

