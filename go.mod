module github.com/joekir/deterministics

go 1.13

require golang.org/x/crypto v0.17.0

// Import local vendor/crypto/rsa module that is pinned pre go 1.11 when they removed determinism
replace crypto/rsa => ./vendor/crypto/rsa
