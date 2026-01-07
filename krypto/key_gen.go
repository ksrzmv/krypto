package krypto

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func GenerateKey(length int) []byte {
	random_data_filepath := "/dev/random"
	random_fd, err := os.Open(random_data_filepath)
	if err != nil {
		panic("could not open random stream")
	}

	key := make([]byte, length)
	n, err := random_fd.Read(key)
	if err != nil {
		random_fd.Close()
		panic(err)
	}
	if n != length {
		random_fd.Close()
		panic("error read from random stream")
	}

	random_fd.Close()
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
