package crypto

import (
	util "github.com/FireStack-Lab/LaksaGo"
	"strings"
	"testing"
)

func TestKeystore_EncryptPrivateKey(t *testing.T) {
	ks := NewDefaultKeystore()
	kv, err := ks.EncryptPrivateKey(util.DecodeHex("24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"), []byte("xiaohuo"), 0)
	if err != nil {
		t.Error(err.Error())
	} else {
		println(kv)
	}
}

func TestKeystore_DecryptPrivateKey(t *testing.T) {
	json := "{\"address\":\"b5c2cdd79c37209c3cb59e04b7c4062a8f5d5271\",\"id\":\"979daaf9-daf1-4002-8656-3cea134c9518\",\"version\":3,\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"26be10cdae0f397bdeead38e7fcc179957dd5e7ef95a1f0f53f37b7ad1355159\",\"kdf\":\"pbkdf2\",\"mac\":\"81d8e60bc08237e4ba154c0b27ad08562821d8c602ee8a492434128de48b66bc\",\"cipherparams\":{\"iv\":\"fc714ad6267c35a2df4cb3f8b8b3cc0d\"},\"kdfparams\":{\"n\":8192,\"c\":262144,\"r\":8,\"p\":1,\"dklen\":32,\"salt\":\"e22ef8a67a59299cee1532b6c6967bdfb0e75ca3c5dff852f9d8daa04683b0c1\"}}}"

	ks := NewDefaultKeystore()
	privateKey, err := ks.DecryptPrivateKey(json, "xiaohuo")
	if err != nil {
		t.Error(err.Error())
	} else {
		if strings.Compare(strings.ToLower(privateKey), "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9") != 0 {
			t.Error("decrypt private key failed")
		}
	}
}
