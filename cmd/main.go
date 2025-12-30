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

	"krypto/krypto"
)

func main() {
	var m krypto.Mode = krypto.Enc
	is_decrypt := flag.Bool("d", false, "enter decryption mode")
	is_key := flag.Bool("k", false, "enter key generation mode")
	read_key_from_file := flag.Bool("K", false, "switch to read key from file './.kr-dek'")
	file_path := os.Args[len(os.Args)-1]
	key_filepath := "./.kr-dek"

	flag.Parse()
	if *is_decrypt && *is_key == true {
		fmt.Println("choose either -d, or -k")
		return
	}
	if *is_decrypt == true {
		m = krypto.Dec
	}
	if *is_key == true {
		m = krypto.Key
	}

	if m == krypto.Key {
		key := krypto.GenerateKey(255)
		binary.Write(os.Stdout, binary.LittleEndian, key)
		return
	}

	data, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	var key []byte
	if *read_key_from_file == false {
		key = krypto.ReadKeyFromTerminal()
	} else {
		key = krypto.ReadKeyFromFile(key_filepath)
	}

	if m == krypto.Enc {
	  encrypted_data := krypto.Encrypt(data, key)
	  enc_file := file_path + ".enc"
	  _ = os.WriteFile(enc_file, encrypted_data, 0666)
	  //binary.Write(encFd, binary.LittleEndian, encrypted_data)
	}

	if m == krypto.Dec {
		decrypted_data := krypto.Decrypt(data, key)
		dec_file := file_path + ".dec"
		_ = os.WriteFile(dec_file, decrypted_data, 0666)
	  //binary.Write(decFd, binary.LittleEndian, decrypted_data)
	}
}
