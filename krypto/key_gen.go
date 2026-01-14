package krypto

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// generate key from /dev/random stream
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

// read key from file
func ReadKeyFromFile(key_filepath string) ([]byte, error) {
	key, err := os.ReadFile(key_filepath)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// read key from terminal
func ReadKeyFromTerminal() ([]byte, error) {
	fmt.Printf("Enter secret key:\n> ")

	// hidden input
	key, err := terminal.ReadPassword(0)
	if err != nil {
		return nil, err
	}

	return key, nil
}
