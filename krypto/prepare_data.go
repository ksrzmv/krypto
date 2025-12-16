package krypto

func alignWord(x uint) uint {
	if x == 0 {
		return 1
	}
	if x % KR_WORD_SIZE_BYTES == 0 {
		return x / KR_WORD_SIZE_BYTES
	} else {
		return x / KR_WORD_SIZE_BYTES + 1
	}
}

func dataToUintArray(data []byte, m Mode) []uint {
	dataLength := uint(len(data))
	wordBlocks := alignWord(dataLength)
	if m == Enc && dataLength % KR_DWORD_SIZE_BYTES == 0 {
		wordBlocks += 2
	} else if m == Enc && wordBlocks % 2 != 0 {
		wordBlocks += 1
	}

	delta := wordBlocks * KR_WORD_SIZE_BYTES - dataLength

	var counter uint

	preparedData := make([]uint, wordBlocks)
	for counter = 0; counter < dataLength; counter++ {
		preparedData[counter/KR_WORD_SIZE_BYTES] += uint(data[counter]) << (56 - (counter % KR_WORD_SIZE_BYTES)*KR_WORD_SIZE_BYTES)
	}

	preparedData[wordBlocks-1] += delta

	return preparedData
}

func dataFromUintArray(data []uint, m Mode) []byte {
	byteDataLength := uint(len(data)) * KR_WORD_SIZE_BYTES
	// TODO: strip trail zeroes when decryption
	if m == Dec {
		byteDataLength -= uint(byte(data[len(data)-1]))
	}
	byteData := make([]byte, byteDataLength)
	for i := 0; uint(i) < byteDataLength; i++ {
		byteData[i] = byte(Rotl(data[i/KR_WORD_SIZE_BYTES], uint(i+1) % KR_WORD_SIZE_BYTES * KR_WORD_SIZE_BYTES))
	}

	return byteData
}

