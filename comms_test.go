package main

import (
	"fmt"
	"testing"
)

func TestComms(t *testing.T) {
	fmt.Println("Generating key and ID...")

	id := GenerateID()
	priv := Generate()
	str := Stringify(priv)
	fmt.Println(str)

	fmt.Println("Uploading...")
	err := PostKey(priv, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Key uploaded")

	fmt.Println("Retrieving key...")
	priv, err = GetKey(id)

	if err != nil {
		panic(err)
	}

	fmt.Println(Stringify(priv))
}

func TestServer(t *testing.T) {
	fmt.Println("Sending the same key twice...")
	priv := Generate()
	id := GenerateID()
	PostKey(priv, id)
	PostKey(priv, id)
}
