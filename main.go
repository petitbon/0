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
	app.Name = "zero"
	app.Usage = "Help me help you!"

	app.Action = func(c *cli.Context) {
		println("Type '0 -h' to summon help.")
	}
	app.Commands = []cli.Command{
		{
			Name:  "hmac-tag",
			Usage: "Tag a --message using a --key. Use '0 hmactag -h' to summont help.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Usage: "HMAC secret key.",
				},
				cli.StringFlag{
					Name:  "message, m",
					Usage: "Message to tag.",
				},
			},

			Action: HMACSign,
		},
		{
			Name:  "hmac-verify",
			Usage: "Verify if a --tag is equal to a --message tagged using a --key. Use '0 hmacverify -h' to summon help.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Usage: "HMAC secret key.",
				},
				cli.StringFlag{
					Name:  "message, m",
					Usage: "Message to verify",
				},
				cli.StringFlag{
					Name:  "tag",
					Usage: "Tag to verify.",
				},
			},
			Action: HMACVerify,
		},
	}

	app.Run(os.Args)
}

func HMACSign(c *cli.Context) {
	if len(c.FlagNames()) == 2 {
		println("tag ", ComputeHmac256(c.String("message"), c.String("key")))
	} else {
		println("Please provide a --message and a --key.")
	}

}

func HMACVerify(c *cli.Context) {
	if len(c.FlagNames()) == 3 {
		println("verified ", VerifyHmac256(c.String("message"), c.String("tag"), c.String("key")))
	} else {
		println("Please provide a --message with a --key and a --tag.")
	}
}

func ComputeHmac256(message string, key string) string {
	k := []byte(key)
	h := hmac.New(sha256.New, k)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func VerifyHmac256(message string, tag string, key string) bool {
	t := []byte(tag)
	k := []byte(key)
	h := hmac.New(sha256.New, k)
	h.Write([]byte(message))
	expectedMAC := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return hmac.Equal(t, []byte(expectedMAC))
}

