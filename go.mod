module github.com/ksrzmv/krypto

go 1.25.1

require (
	golang.org/x/crypto v0.46.0
	pkg/krypto v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/term v0.38.0 // indirect
)

replace pkg/krypto => ./pkg/krypto
