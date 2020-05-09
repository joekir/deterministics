package deterministics

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/binary"
	"encoding/pem"
	"golang.org/x/crypto/pbkdf2"
	"math/rand"
	"os"
)

func DeriveKeys(passphrase, privateKeyFile, pubKeyFile string) error {
	// yes this sucks for now
	// perhaps in future some per instance salt?
	salt := []byte("")
	dk := pbkdf2.Key([]byte(passphrase), salt, 100000, 64, sha256.New)

	seed := int64(binary.LittleEndian.Uint64(dk))
	r := rand.New(rand.NewSource(seed))

	key, err := rsa.GenerateKey(r, 4096)
	if err != nil {
		return err
	}

	err = savePrivatePEMKey(privateKeyFile, key)
	if err != nil {
		return err
	}

	err = savePublicPEMKey(pubKeyFile, key.PublicKey)
	return err
}

func savePrivatePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	return err
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	return err
}
