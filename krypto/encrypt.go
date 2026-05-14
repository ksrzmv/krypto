// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package krypto

func Encrypt(data []byte, key []byte) ([]byte, error) {
	// convert byte array to uint blocks
	preparedData, err := dataToUintArray(data, Enc)
	if err != nil {
		return nil, err
	}

	// create expansion key array
	S := keyExpansion(key)

	// enrich data with initialization vector (IV)

	// generate dword of random data for IV block
	iv, err := GenerateKey(KR_DWORD_SIZE_BYTES)
	if err != nil {
		return nil, err
	}

	// length of enriched data = length of actual data + 2 iv blocks
	initializationVectorLength := len(preparedData) + 2
	initializationVectorData := make([]uint, initializationVectorLength)
	for i := 0; i < len(iv)-KR_WORD_SIZE_BYTES; i++ {
		initializationVectorData[0] += uint(iv[i] << (56 - i*KR_WORD_SIZE_BYTES))
		initializationVectorData[1] += uint(iv[i+KR_WORD_SIZE_BYTES] << (56 - i*KR_WORD_SIZE_BYTES))
	}
	for idx, val := range preparedData {
		initializationVectorData[idx+2] = val
	}
	preparedData = initializationVectorData
	// -----

	// range over all two-block batches
	for i := 0; i < len(preparedData)-1; i += 2 {
		A := preparedData[i]
		B := preparedData[i+1]

		if i > 1 {
			A ^= preparedData[i-2]
			B ^= preparedData[i-1]
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

		preparedData[i] = A
		preparedData[i+1] = B
	}

	// convert uint blocks to byte array
	result, err := dataFromUintArray(preparedData, Enc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Decrypt(data []byte, key []byte) ([]byte, error) {
	// convert byte array to uint blocks
	preparedData, err := dataToUintArray(data, Dec)
	if err != nil {
		return nil, err
	}

	// create expansion key array
	S := keyExpansion(key)

	// CBC mode feedback; initialized as 0 for first XOR
	cbcFeedbackA := uint(0)
	cbcFeedbackB := uint(0)
	var previousBlockA, previousBlockB uint

	// range over all two-block batches
	for i := 0; i < len(preparedData)-1; i += 2 {
		A := preparedData[i]
		B := preparedData[i+1]

		previousBlockA = A
		previousBlockB = B

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

		B ^= cbcFeedbackB
		A ^= cbcFeedbackA

		// feedback for next round is previous encrypted two-blocks batch
		cbcFeedbackA = previousBlockA
		cbcFeedbackB = previousBlockB

		preparedData[i+1] = B
		preparedData[i] = A
	}

	// convert uint blocks to byte array
	result, err := dataFromUintArray(preparedData[2:], Dec)
	if err != nil {
		return nil, err
	}

	return result, nil
}
