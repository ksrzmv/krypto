package main

import (
	"fmt"
	"testing"

	"github.com/ksrzmv/krypto/krypto"
)

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
