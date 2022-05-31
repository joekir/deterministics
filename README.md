### Deterministics

 *A tool to generate a deterministic RSA keypair from a passphrase.*

 Caveat: this isn't for ideal world scenarios.

### Why?

Sometimes it's required to have symmetric key material on disk that
is derived from user input. 

In some scenarios the key is only used to encrypt, and the decryption is 
done elsewhere. Hence asymmetric keys make sense.
 
### The approach

This was inspired by the response to a question by Thomas Pornin:
https://crypto.stackexchange.com/a/1665/7763

### Testing

```
$ go build
$ ./deterministics mypassphrase
KeyGen Complete.
$ go test -v ./...
=== RUN   TestRoundTripWithGeneratedKeys
    outputs_test.go:15: reading RSA key from outdir/test.priv
    outputs_test.go:22: RSA PRIVATE KEY
    outputs_test.go:48: PUBLIC KEY
--- PASS: TestRoundTripWithGeneratedKeys (0.01s)
PASS
ok      github.com/joekir/deterministics        0.017s
=== RUN   TestSameKeyGeneratedEachTime
    deterministics_test.go:18: /tmp/deterministics_tests004736468
--- PASS: TestSameKeyGeneratedEachTime (1.11s)
PASS
ok      github.com/joekir/deterministics/lib    1.115s
```

