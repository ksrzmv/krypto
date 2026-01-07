package krypto

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func GenerateKey(length int) ([]byte, error) {
	random_data_filepath := "/dev/random1"
	random_fd, err := os.Open(random_data_filepath)
	if err != nil {
		fmt.Println("could not open random stream\nerror:", err)
		return nil, err
	}

	key := make([]byte, length)
	n, err := random_fd.Read(key)

	defer random_fd.Close()
	if err != nil {
		fmt.Println("could not read from random fd\nerror:", err)
		return nil, err
	}
	if n != length {
		fmt.Println("invalid read length from random fd\nerror:", err)
		return nil, err
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
