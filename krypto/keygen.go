package krypto

import (
	"os"
)

func GenerateKey(length int) []byte {
	randomDataFilePath := "/dev/random"
	randomFd, err := os.Open(randomDataFilePath)
	if err != nil {
		panic("could not open random stream")
	}

	key := make([]byte, length)
	n, err := randomFd.Read(key)
	if err != nil {
		panic(err)
	}
	if n != length {
		panic("error read from random stream")
	}

	return key
}
