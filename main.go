package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "zero cli"
	app.Usage = "Help me help you!"

	app.Action = func(c *cli.Context) {
		println("Type '0 -h' to summon the help.")
	}
	app.Commands = []cli.Command{
		{
			Name:  "hs",
			Usage: "sign or verify messages base64 hash using HMAC SHA256. Use '0 hmac -h' to get help.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "secret",
					Usage: "HMAC secret",
				},
			},
			Action: HMACSign,
		},
		{
			Name:  "hv",
			Usage: "sign or verify messages base64 hash using HMAC SHA256. Use '0 hmac -h' to get help.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "s",
					Usage: "HMAC secret",
				},
			},
			Action: HMACVerify,
		},
	}

	app.Run(os.Args)
}

func HMACSign(c *cli.Context) {
	if len(c.Args()) > 0 {
		println("message signed: ", c.Args()[0], c.String("secret"))
	} else {
		println("Please provide a message to sign.")
	}
}

func HMACVerify(c *cli.Context) {
	if len(c.Args()) > 0 {
		println("message: ", c.Args()[0], " secret: ", c.String("secret"), " signedmessage: ", ComputeHmac256(c.Args()[0], c.String("secret")))
	} else {
		println("Please provide a message to sign.")
	}
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

/*
..





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
	flagSign   = flag.String("sign", "ijnenwyh", "")
	flagVerify = flag.String("verify", "ijnenwyh", "")
)

var usage = `Usage: 0 [options...] message

Options:
  -sign <secret> Creates a base64 hash using HMAC SHA256
  -verify <secret> Verify HMAC signed message
`

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	flag.Parse()
	var sign = *flagSign

	fmt.Println(sign)

	fmt.Println(ComputeHmac256("hello", sign))
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
*/
