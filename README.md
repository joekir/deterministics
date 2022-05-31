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
$ go test
PASS
ok      github.com/joekir/deterministics        0.013s
```

