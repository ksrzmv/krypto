package krypto

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// GenerateKey - generates key from /dev/random data stream
func GenerateKey(length int) ([]byte, error) {
	randomDataFilepath := "/dev/random"
	randomFileDescriptor, err := os.Open(randomDataFilepath)
	if err != nil {
		return nil, err
	}

	key := make([]byte, length)
	n, err := randomFileDescriptor.Read(key)

	defer func(randomFileDescriptor *os.File) {
		err := randomFileDescriptor.Close()
		if err != nil {
			fmt.Println("error close random file descriptor: ", err)
		}
	}(randomFileDescriptor)

	if err != nil {
		return nil, err
	}
	if n != length {
		return nil, errors.New("invalid read length from random stream")
	}

	return key, nil
}

// ReadKeyFromFile - reads key from file
func ReadKeyFromFile(keyFilepath string) ([]byte, error) {
	key, err := os.ReadFile(keyFilepath)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// ReadKeyFromTerminal - reads key from user input
func ReadKeyFromTerminal() ([]byte, error) {
	fmt.Printf("Enter secret key:\n> ")

	// hidden input
	key, err := terminal.ReadPassword(0)
	if err != nil {
		return nil, err
	}

	return key, nil
}
