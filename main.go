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
var (
	check        bool
	testStrings  string
	hash         string
	randSeqBytes []byte
)

func main() {
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

	fmt.Printf("String     : %s\n", testStrings)
	fmt.Printf("SHA256 hash: %s\n", hash)
}

// Generates a random string of length n
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(b)
}
