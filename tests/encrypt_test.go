package krypto_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/ksrzmv/krypto/krypto"
)

func slices_equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func TestEncrypt(t *testing.T) {
	max_data_len := 1000000
	max_key_len := 256
	tests_count := 1000

	var wg sync.WaitGroup

	for i := range tests_count {
		data_length := rand.Intn(max_data_len)
		data, err := krypto.GenerateKey(data_length)
		if err != nil {
			t.Errorf("error generate data of size: %d", data_length)
		}
		key_length := rand.Intn(max_key_len)
		key, err := krypto.GenerateKey(key_length)
		if err != nil {
			t.Errorf("error generate key of size: %d", key_length)
		}
		name := fmt.Sprintf("%d data length: %d, key length: %d", i,  data_length, key_length)
		wg.Add(1)
		go t.Run(name, func(t *testing.T) {
			defer wg.Done()
			encrypted_data, err := krypto.Encrypt(data, key)
			if err != nil {
				t.Errorf("error encrypt data. data len: %d, key len: %d", data_length, key_length)
			}
			decrypted_data, err := krypto.Decrypt(encrypted_data, key)
			if err != nil {
				t.Errorf("error decrypt data. data len: %d, key len: %d", data_length, key_length)
			}

			if !slices_equal(decrypted_data, data) {
				t.Errorf("decrypted data doesn't match original data")
			}
		})
	}

	wg.Wait()
}

func BenchmarkEncrypt(b *testing.B) {
	data_lengths := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	key_length := 255
	key, _ := krypto.GenerateKey(key_length)

	for _, ln := range data_lengths {
		data, _ := krypto.GenerateKey(ln)
		name := fmt.Sprintf("Data Length: %d", ln)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				_, _ = krypto.Encrypt(data, key)
			}
		})
	}
}
