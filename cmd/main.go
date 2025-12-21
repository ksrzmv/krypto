// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"krypto/krypto"
)

func main() {
	message := "oleg"

	data := []byte(message)
	key := []byte("oleg")
	encryptedData := krypto.Encrypt(data, key)
	binary.Write(os.Stdout, binary.LittleEndian, encryptedData)
	fmt.Println()
	decryptedData := krypto.Decrypt(encryptedData, key)
	binary.Write(os.Stdout, binary.LittleEndian, decryptedData)
}
