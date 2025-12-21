// rc5-w/r/b
// w - word size (= 32 or 64 bits; depends on architecture due to speed purposes)
// r - # of rounds (= 255)
// b - key length in bytes (up to 256)

package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"krypto/krypto"
)

func main() {
	var m krypto.Mode = krypto.Enc
	isDecrypt := flag.Bool("d", false, "enter decryption mode")
	isKey := flag.Bool("k", false, "enter key generation mode")
	readKeyFromFile := flag.Bool("K", false, "switch to read key from file './.xch-key'")
	filePath := os.Args[len(os.Args)-1]
	keyFilePath := "./.xch-key"

	flag.Parse()
	if *isDecrypt && *isKey == true {
		fmt.Println("choose either -d, or -k")
		return
	}
	if *isDecrypt == true {
		m = krypto.Dec
	}
	if *isKey == true {
		m = krypto.Key
	}

	if m == krypto.Key {
		key := krypto.GenerateKey(255)
		binary.Write(os.Stdout, binary.LittleEndian, key)
		return
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var key []byte
	if *readKeyFromFile == false {
		fmt.Printf("Enter secret key:\n> ")
		key, err = terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}
	} else {
		key, err = os.ReadFile(keyFilePath)
		if err != nil {
			panic(err)
		}
	}

	if m == krypto.Enc {
	  encryptedData := krypto.Encrypt(data, key)
	  encFile := filePath + ".enc"
	  _ = os.WriteFile(encFile, encryptedData, 0666)
	  //binary.Write(encFd, binary.LittleEndian, encryptedData)
	}

	if m == krypto.Dec {
		decryptedData := krypto.Decrypt(data, key)
		decFile := filePath + ".dec"
		_ = os.WriteFile(decFile, decryptedData, 0666)
	}
	//binary.Write(decFd, binary.LittleEndian, decryptedData)
}
