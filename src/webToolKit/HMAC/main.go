package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getHash("YOO YOO")
	fmt.Println(c)
	c = getHash("yoo yoo")
	fmt.Println(c)
}

func getHash(s string) string {
	h := hmac.New(sha256.New, []byte("ourKey"))
	io.WriteString(h, s) //write string into the hash
	return fmt.Sprintf("%x", h.Sum(nil))
}
