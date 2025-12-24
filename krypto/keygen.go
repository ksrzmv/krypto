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
		randomFd.Close()
		panic(err)
	}
	if n != length {
		randomFd.Close()
		panic("error read from random stream")
	}

	randomFd.Close()
	return key
}
