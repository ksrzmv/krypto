package krypto_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/ksrzmv/krypto/krypto"
)

func slicesEqual[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func TestEncrypt(t *testing.T) {
	maxDataLength := 1000000
	maxKeyLength := 256
	testsCount := 1000

	var wg sync.WaitGroup

	for i := range testsCount {
		dataLength := rand.Intn(maxDataLength)
		data, err := krypto.GenerateKey(dataLength)
		if err != nil {
			t.Errorf("error generate data of size: %d", dataLength)
		}
		keyLength := rand.Intn(maxKeyLength)
		key, err := krypto.GenerateKey(keyLength)
		if err != nil {
			t.Errorf("error generate key of size: %d", keyLength)
		}
		name := fmt.Sprintf("%d data length: %d, key length: %d", i, dataLength, keyLength)
		wg.Add(1)
		go t.Run(name, func(t *testing.T) {
			defer wg.Done()
			encryptedData, err := krypto.Encrypt(data, key)
			if err != nil {
				t.Errorf("error encrypt data. data len: %d, key len: %d", dataLength, keyLength)
			}
			decryptedData, err := krypto.Decrypt(encryptedData, key)
			if err != nil {
				t.Errorf("error decrypt data. data len: %d, key len: %d", dataLength, keyLength)
			}

			if !slicesEqual(decryptedData, data) {
				t.Errorf("decrypted data doesn't match original data")
			}
		})
	}

	wg.Wait()
}

func BenchmarkEncrypt(b *testing.B) {
	dataLengths := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	keyLength := 255
	key, _ := krypto.GenerateKey(keyLength)

	for _, ln := range dataLengths {
		data, _ := krypto.GenerateKey(ln)
		name := fmt.Sprintf("Data Length: %d", ln)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				_, _ = krypto.Encrypt(data, key)
			}
		})
	}
}
