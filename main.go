package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var alphaNumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() {
	var check bool
	var testStrings string
	var hash string
	var randSeqBytes []byte

	// loops through the hash256 strings until it finds a match wioth b00da
	for !check {
		rand.Seed(time.Now().UnixNano())
		testStrings = randSeq(3) // 3 is a very short and sweet number to produce a solution
		randSeqBytes = []byte(testStrings)
		h := sha256.New()
		h.Write(randSeqBytes)
		hash = hex.EncodeToString(h.Sum(nil))
		check = strings.Contains(hash, "b00da")
	}

	fmt.Println(testStrings)
	fmt.Println(hash)
}

// Generates a random string using
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(b)
}