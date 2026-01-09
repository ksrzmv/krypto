package krypto

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func GenerateKey(length int) ([]byte, error) {
	random_data_filepath := "/dev/random"
	random_fd, err := os.Open(random_data_filepath)
	if err != nil {
		return nil, err
	}

	key := make([]byte, length)
	n, err := random_fd.Read(key)

	defer random_fd.Close()
	if err != nil {
		return nil, err
	}
	if n != length {
		return nil, errors.New("invalid read length from random stream")
	}

	return key, nil
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
