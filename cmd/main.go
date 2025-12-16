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
	//message := "10513u5[ewjt'au9-t215u82315rh5 23;15h23i5=12 3uy5h02358h21p58 h2p5h23185h23[58h 21-5 8u-5 h132-5b "
	message := "oleg"
	//message := "oleg"

	data := []byte(message)
	key := []byte("oleg")
	encryptedData := krypto.Encrypt(data, key)
	binary.Write(os.Stdout, binary.LittleEndian, encryptedData)
	fmt.Println()
	decryptedData := krypto.Decrypt(encryptedData, key)
	binary.Write(os.Stdout, binary.LittleEndian, decryptedData)
}
