package krypto

import (
	"errors"
	"fmt"
)

// calculates how much word-size blocks needed to store x bytes
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

// converts byte array to uint array
// example. assume we have byte array ["0x5a", "0x11", "0x91", "0xab", "0xfc"]
// so this array will be converted to uint value 0x5a1191abfc000000
func dataToUintArray(data []byte, m Mode) ([]uint, error) {
	data_length := uint(len(data))
	// checks if the input data is properly aligned for decryption
	if m == Dec && data_length % KR_DWORD_SIZE_BYTES != 0 {
		return nil, fmt.Errorf("integrity check error: %w", errors.New("file size for decryption must be divisible by 16 bytes"))
	}

	// additional blocks to store delta when encryption
	word_blocks := alignWord(data_length)
	if m == Enc && data_length % KR_DWORD_SIZE_BYTES == 0 {
		word_blocks += 2
	} else if m == Enc && word_blocks % 2 != 0 {
		word_blocks += 1
	}

	// number of bytes which needed to be added to data for encryption
	delta := word_blocks * KR_WORD_SIZE_BYTES - data_length

	var counter uint

	// all align bytes are zeroes (allocated by make) except last one, which will be delta,
	// indicates the number of such align bytes
	prepared_data := make([]uint, word_blocks)
	for counter = 0; counter < data_length; counter++ {
		prepared_data[counter/KR_WORD_SIZE_BYTES] += uint(data[counter]) << (KR_WORD_SIZE - 8 - (counter % KR_WORD_SIZE_BYTES)*KR_WORD_SIZE_BYTES)
	}

	prepared_data[word_blocks-1] += delta

	return prepared_data, nil
}

// converts uint array to byte array
func dataFromUintArray(data []uint, m Mode) ([]byte, error) {
	byte_data_length := uint(len(data)) * KR_WORD_SIZE_BYTES
	// remove align bytes when performed decryption.
	// if delta (last byte) is more than DWORD size, returns error.
	// there's could be false-positive errors when DEK was incorrect:
	// if decryption was invalid (due to incorrect DEK), but last byte is smaller than DWORD size,
	// the error will not be returned, but output data will be a mess
	if m == Dec {
		delta := uint(byte(data[len(data)-1]))
		if delta > KR_DWORD_SIZE_BYTES {
			return nil, errors.New("invalid decryption. possibly DEK is incorrect")
		}
		byte_data_length -= delta
	}
	byte_data := make([]byte, byte_data_length)
	for i := 0; uint(i) < byte_data_length; i++ {
		byte_data[i] = byte(Rotl(data[i/KR_WORD_SIZE_BYTES], uint(i+1) % KR_WORD_SIZE_BYTES * KR_WORD_SIZE_BYTES))
	}

	return byte_data, nil
}

