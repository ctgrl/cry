package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"os"

	"crypto/sha256"
)

func decrypt(file string, priv *rsa.PrivateKey) {
	inFile, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer inFile.Close()

	outFile, err := os.OpenFile(file[:len(file)-len(LockedExtension)], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	header := make([]byte, EncryptedHeaderSize)
	_, err = io.ReadFull(inFile, header)
	if err != nil {
		panic(err)
	}

	label := []byte("")

	header, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, header, label)

	if err != nil {
		panic(err)
	}

	key := header[:KeySize]
	iv := header[KeySize : KeySize+aes.BlockSize]

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	reader := &cipher.StreamReader{S: stream, R: inFile}

	_, err = io.Copy(outFile, reader)
	if err != nil {
		panic(err)
	}
}
