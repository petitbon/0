package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

var (
	flagCommand = flag.String("sign", "secret", "")
)

var usage = `Usage: 0 [options...] <url>

Options:
  -h  Creates a base64 hash using HMAC SHA256 
`

func main() {

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	flag.Parse()
	var command = *flagCommand

	fmt.Println(command)

	fmt.Println(ComputeHmac256("hello", command))
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func usageAndExit(message string) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
