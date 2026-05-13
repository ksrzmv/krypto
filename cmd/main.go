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

	"github.com/ksrzmv/krypto/krypto"
)

func main() {
	m := krypto.Enc

	isDecrypt := flag.Bool("d", false, "decryption mode")
	isKey := flag.Bool("k", false, "keygen mode. generates 255 bytes key from /dev/random, outputs to stdout")
	readKeyFromFile := flag.Bool("K", false, "switch to read key from file './.kr-dek'")
	filePath := os.Args[len(os.Args)-1]
	keyFilepath := "./.kr-dek"

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
		key, err := krypto.GenerateKey(255)
		if err != nil {
			fmt.Println("could not generate key\nerror:", err)
			return
		}
		err = binary.Write(os.Stdout, binary.LittleEndian, key)
		if err != nil {
			fmt.Println("error write binary key to stdout\nerror:", err)
			return
		}
		return
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("cannot read input file, exit\nerror:", err)
		return
	}

	var key []byte
	if *readKeyFromFile == false {
		key, err = krypto.ReadKeyFromTerminal()
		if err != nil {
			fmt.Println("cannot read key from terminal, exit\nerror:", err)
			return
		}
	} else {
		key, err = krypto.ReadKeyFromFile(keyFilepath)
		if err != nil {
			fmt.Println("cannot read key file, exit\nerror:", err)
			return
		}
	}

	if m == krypto.Enc {
		encryptedData, err := krypto.Encrypt(data, key)
		if err != nil {
			fmt.Println(err)
			return
		}
		encFile := filePath + ".enc"
		err = os.WriteFile(encFile, encryptedData, 0666)
		if err != nil {
			fmt.Println("failed to write encrypted data to file\nerror:", err)
		}
		//binary.Write(encFd, binary.LittleEndian, encrypted_data)
	}

	if m == krypto.Dec {
		decryptedData, err := krypto.Decrypt(data, key)
		if err != nil {
			fmt.Println(err)
			return
		}
		decFile := filePath + ".dec"
		_ = os.WriteFile(decFile, decryptedData, 0666)
		//binary.Write(decFd, binary.LittleEndian, decrypted_data)
	}
}
