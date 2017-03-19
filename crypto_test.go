package main

import (
	"fmt"
	"testing"
)

func TestCrypto(t *testing.T) {
	file := "test.docx"
	priv := Generate()

	fmt.Println("Encrypting test file...")
	encrypt(file, priv)
	fmt.Println("Decrypting test file...")
	decrypt(file+LockedExtension, priv)
}
