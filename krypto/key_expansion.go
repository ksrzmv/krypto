package krypto

func keyExpansion(key []byte) []uint {
	keyLength := len(key)
	byteWordsToFillKey := alignWord(uint(keyLength))
	L := make([]uint, byteWordsToFillKey)
	for i := keyLength-1; i > -1; i-- {
		L[i/KR_WORD_SIZE_BYTES] = (Rotl(L[i/KR_WORD_SIZE_BYTES], 8) + uint(key[i])) & KR_MODULUS
	}

	sLength := uint(2*(KR_ROUNDS + 1))
	S := make([]uint, sLength)
	S[0] = P
	for i := 1; uint(i) < sLength; i++ {
		S[i] = (S[i-1] + Q) & KR_MODULUS
	}

	i := uint(0)
	j := uint(0)
	A := uint(0)
	B := uint(0)
	var counter uint
	for counter = 0; counter < 3 * max(sLength, byteWordsToFillKey); counter++ {
		S[i] = Rotl((S[i] + A + B) & KR_MODULUS, 3)
		L[j] = Rotl((L[j] + A + B) & KR_MODULUS, (A + B) & KR_MODULUS)
		A = S[i]
		B = L[j]
		i = (i + 1) % sLength
		j = (j + 1) % byteWordsToFillKey
	}

	return S
}

