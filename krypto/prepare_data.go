package krypto

import (
	"errors"
)

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

func dataToUintArray(data []byte, m Mode) ([]uint, error) {
	data_length := uint(len(data))
	if m == Dec && data_length % KR_DWORD_SIZE_BYTES != 0 {
		return nil, errors.New("integrity check error. file size for decryption must be divisible by 16 bytes")
	}
	word_blocks := alignWord(data_length)
	if m == Enc && data_length % KR_DWORD_SIZE_BYTES == 0 {
		word_blocks += 2
	} else if m == Enc && word_blocks % 2 != 0 {
		word_blocks += 1
	}

	delta := word_blocks * KR_WORD_SIZE_BYTES - data_length

	var counter uint

	prepared_data := make([]uint, word_blocks)
	for counter = 0; counter < data_length; counter++ {
		prepared_data[counter/KR_WORD_SIZE_BYTES] += uint(data[counter]) << (KR_WORD_SIZE - 8 - (counter % KR_WORD_SIZE_BYTES)*KR_WORD_SIZE_BYTES)
	}

	prepared_data[word_blocks-1] += delta

	return prepared_data, nil
}

func dataFromUintArray(data []uint, m Mode) ([]byte, error) {
	byte_data_length := uint(len(data)) * KR_WORD_SIZE_BYTES
	if m == Dec {
		delta := uint(byte(data[len(data)-1]))
		if delta > KR_DWORD_SIZE_BYTES {
			return nil, errors.New("invalid decryption")
		}
		byte_data_length -= delta
	}
	byte_data := make([]byte, byte_data_length)
	for i := 0; uint(i) < byte_data_length; i++ {
		byte_data[i] = byte(Rotl(data[i/KR_WORD_SIZE_BYTES], uint(i+1) % KR_WORD_SIZE_BYTES * KR_WORD_SIZE_BYTES))
	}

	return byte_data, nil
}

