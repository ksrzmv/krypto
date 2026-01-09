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

	is_decrypt := flag.Bool("d", false, "decryption mode")
	is_key := flag.Bool("k", false, "keygen mode. generates 255 bytes key from /dev/random, outputs to stdout")
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
		key, err := krypto.GenerateKey(255)
		if err != nil {
			fmt.Println("could not generate key\nerror:", err)
			return
		}
		binary.Write(os.Stdout, binary.LittleEndian, key)
		return
	}

	data, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("cannot read input file, exit\nerror:", err)
		return
	}

	var key []byte
	if *read_key_from_file == false {
		key = krypto.ReadKeyFromTerminal()
	} else {
		key, err = krypto.ReadKeyFromFile(key_filepath)
		if err != nil {
			fmt.Println("cannot read key file, exit\nerror:", err)
			return
		}
	}

	if m == krypto.Enc {
	  encrypted_data, err := krypto.Encrypt(data, key)
		if err != nil {
			fmt.Println(err)
			return
		}
	  enc_file := file_path + ".enc"
	  _ = os.WriteFile(enc_file, encrypted_data, 0666)
	  //binary.Write(encFd, binary.LittleEndian, encrypted_data)
	}

	if m == krypto.Dec {
		decrypted_data, err := krypto.Decrypt(data, key)
		if err != nil {
			fmt.Println(err)
			return
		}
		dec_file := file_path + ".dec"
		_ = os.WriteFile(dec_file, decrypted_data, 0666)
	  //binary.Write(decFd, binary.LittleEndian, decrypted_data)
	}
}
