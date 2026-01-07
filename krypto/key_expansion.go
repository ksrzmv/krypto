package krypto

func keyExpansion(key []byte) []uint {
	key_length := len(key)
	byte_words_to_fill_key := int(alignWord(uint(key_length)))
	L := make([]uint, byte_words_to_fill_key)
	for i := key_length-1; i > -1; i-- {
		L[i/KR_WORD_SIZE_BYTES] = (Rotl(L[i/KR_WORD_SIZE_BYTES], 8) + uint(key[i])) & KR_MODULUS
	}

	s_length := 2*(KR_ROUNDS + 1)
	S := make([]uint, s_length)
	S[0] = P
	for i := 1; i < s_length; i++ {
		S[i] = (S[i-1] + Q) & KR_MODULUS
	}

	i := 0
	j := 0
	A := uint(0)
	B := uint(0)
	for counter := 0; counter < 3 * max(s_length, byte_words_to_fill_key); counter++ {
		S[i] = Rotl((S[i] + A + B) & KR_MODULUS, 3)
		L[j] = Rotl((L[j] + A + B) & KR_MODULUS, (A + B) & KR_MODULUS)
		A = S[i]
		B = L[j]
		i = (i + 1) % s_length
		j = (j + 1) % byte_words_to_fill_key
	}

	return S
}

