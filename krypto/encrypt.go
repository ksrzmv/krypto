// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package krypto

func Encrypt(data []byte, key []byte) []byte {
	prepData := dataToUintArray(data, Enc)
	expandedKey := keyExpansion(key)

	for i := 0; i < len(prepData) - 1; i += 2 {
		prepData[i] += expandedKey[0]
		prepData[i+1] += expandedKey[1]
		for j := 1; j < KR_ROUNDS; j += 1 {
			prepData[i] ^= prepData[i+1]
			prepData[i] = Rotl(prepData[i], prepData[i+1] & KR_MODULUS)
			prepData[i] += expandedKey[2*j]
			prepData[i+1] ^= prepData[i]
			prepData[i+1] = Rotl(prepData[i+1], prepData[i] & KR_MODULUS)
			prepData[i+1] += expandedKey[2*j+1]
			// debugging
		  //fmt.Printf("block: %d round: %d A: %016x\n", i, j, prepData[i])
		  //fmt.Printf("block: %d round: %d B: %016x\n", i+1, j, prepData[i+1])
		}
	}

	return dataFromUintArray(prepData, Enc)
}

func Decrypt(data []byte, key []byte) []byte {
	prepData := dataToUintArray(data, Dec)
	expandedKey := keyExpansion(key)

	for i := 0; i < len(prepData) - 1; i += 2 {
		for j := KR_ROUNDS-1; j > 0; j -= 1 {
			prepData[i+1] -= expandedKey[2*j+1]
			prepData[i+1] = Rotr(prepData[i+1], prepData[i] & KR_MODULUS)
			prepData[i+1] ^= prepData[i]
			prepData[i] -= expandedKey[2*j]
			prepData[i] = Rotr(prepData[i], prepData[i+1] & KR_MODULUS)
			prepData[i] ^= prepData[i+1]
		}
		prepData[i+1] -= expandedKey[1]
		prepData[i] -= expandedKey[0]
	}

	return dataFromUintArray(prepData, Dec)
}

