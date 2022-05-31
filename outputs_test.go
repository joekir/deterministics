package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"testing"
)

func TestRoundTripWithGeneratedKeys(t *testing.T) {
	t.Log("reading RSA key from outdir/test.priv")
	privateKeyPemBytes, err := ioutil.ReadFile("outdir/test.priv")
	if err != nil {
		t.Error(err)
	}

	block, _ := pem.Decode(privateKeyPemBytes)
	t.Log(block.Type)

	if block == nil || block.Type != "RSA PRIVATE KEY" {
		t.Error("not a private key")
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Error(err)
	}

	h := sha256.New()
	h.Write([]byte("hello world"))
	digest := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, digest)
	if err != nil {
		t.Error(err)
	}

	publicKeyPemBytes, err := ioutil.ReadFile("outdir/test.pub")
	if err != nil {
		t.Error(err)
	}

	block, _ = pem.Decode(publicKeyPemBytes)
	t.Log(block.Type)

	if block == nil || block.Type != "PUBLIC KEY" {
		t.Error("not a public key")
	}

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		t.Error(err)
	}

	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, digest, signature)
	if err != nil {
		t.Error(err)
	}

	signature[0] = 0x22
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, digest, signature)
	if err == nil {
		t.Error("Verification should have failed with a bogus signature!")
	}

	if err.Error() != "crypto/rsa: verification error" {
		t.Errorf("Unexpected error message encountered: '%s'\n", err.Error())
	}
}
