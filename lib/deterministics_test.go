package deterministics

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSameKeyGeneratedEachTime(t *testing.T) {
	testDir, err := ioutil.TempDir("/tmp", "deterministics_tests")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testDir)

	t.Log(testDir)

	privPath1 := filepath.Join(testDir, "priv1.pem")
	privPath2 := filepath.Join(testDir, "priv2.pem")
	pubPath1 := filepath.Join(testDir, "pub1.pem")
	pubPath2 := filepath.Join(testDir, "pub2.pem")

	DeriveKeys("coolpassword", privPath1, pubPath1)
	DeriveKeys("coolpassword", privPath2, pubPath2)

	pub1, err := ioutil.ReadFile(pubPath1)
	if err != nil {
		t.Error(err)
	}

	pub2, err := ioutil.ReadFile(pubPath2)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(pub1, pub2) {
		t.Error("Generated deterministic public key files (with same seed phrase) differ!")
	}

	priv1, err := ioutil.ReadFile(privPath1)
	if err != nil {
		t.Error(err)
	}

	priv2, err := ioutil.ReadFile(privPath2)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(priv1, priv2) {
		t.Error("Generated deterministic private key files (with same seed phrase) differ!")
	}

}
