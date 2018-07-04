package uaes

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	var aeskey = []byte("1234567890123456")
	pass := []byte("this is password")
	xpass, err := AesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("Encrypt:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypt:%s\n", tpass)
}
