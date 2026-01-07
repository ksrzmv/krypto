package krypto

const (
	KR_ROUNDS						= 255
	// explaination of KR_WORD_SIZE calculation
	// ^uint(0) - uint value with all bits set (bit inversion; e.g. ^(0b0101) = 0b1010)
	// '>> 63' will be 0 for 32-bit arch and 1 for 64-bit arch
	// '32 <<' : for 32-bit it will result in 32 << 0 = 32, for 64-bit: 32 << 1 = 64
	KR_WORD_SIZE 				= 32 << (^uint(0) >> 63)
	KR_WORD_SIZE_BYTES 	= KR_WORD_SIZE / 8
	KR_DWORD_SIZE_BYTES	= 2 * KR_WORD_SIZE_BYTES
	KR_MODULUS					= (1 << KR_WORD_SIZE) - 1

	// rc5 magic constants. see Rivest's whitepaper for explaination.
	P 									= 0xb7e151628aed2a6b >> (64 - KR_WORD_SIZE)
	Q										= 0x9e3779b97f4a7c15 >> (64 - KR_WORD_SIZE)

)

type Mode byte

const (
	Enc Mode = 0
	Dec Mode = 1
	Key Mode = 2
)

