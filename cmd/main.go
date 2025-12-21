// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"krypto/krypto"
)

func main() {
	var m krypto.Mode = krypto.Enc
	isDecrypt := flag.Bool("m", false, "enter decryption mode")
	filePath := flag.String("file", "", "file encrypt to/decrypt from")
	flag.Parse()
	if *isDecrypt == true {
		m = krypto.Dec
	}
	data, err := os.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Enter secret key:\n> ")
	key, err := terminal.ReadPassword(0)
	if err != nil {
		panic(err)
	}

	if m == krypto.Enc {
	  encryptedData := krypto.Encrypt(data, key)
	  encFile := *filePath + ".enc"
	  _ = os.WriteFile(encFile, encryptedData, 0666)
	  //binary.Write(encFd, binary.LittleEndian, encryptedData)
	}

	if m == krypto.Dec {
		decryptedData := krypto.Decrypt(data, key)
		decFile := *filePath + ".dec"
		_ = os.WriteFile(decFile, decryptedData, 0666)
	}
	//binary.Write(decFd, binary.LittleEndian, decryptedData)
}
