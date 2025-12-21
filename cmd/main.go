// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"krypto/krypto"
)

func main() {
	message := "oleg"

	data := []byte(message)
	fmt.Printf("Enter secret key:\n> ")
	key, err := terminal.ReadPassword(0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n")
	encryptedData := krypto.Encrypt(data, key)
	binary.Write(os.Stdout, binary.LittleEndian, encryptedData)
	fmt.Printf("\n")
	decryptedData := krypto.Decrypt(encryptedData, key)
	binary.Write(os.Stdout, binary.LittleEndian, decryptedData)
}
