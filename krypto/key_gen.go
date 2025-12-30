package krypto

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
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

func ReadKeyFromFile(key_filepath string) ([]byte, error) {
	key, err := os.ReadFile(key_filepath)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func ReadKeyFromTerminal() []byte {
	fmt.Printf("Enter secret key:\n> ")
	key, err := terminal.ReadPassword(0)
	if err != nil {
		panic(err)
	}

	return key
}
