package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
	//"github.com/petitbon/zerolib"
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
			Name:  "curl",
			Usage: "Curl.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url, u",
					Usage: "--url.",
				},
			},

			Action: SimpleCurl,
		},

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

// HMAC FUNCTIONS //////////////////////////////////////////

func HMACSign(c *cli.Context) {
	if len(c.FlagNames()) == 2 {
		println("tag ", ComputeHmac256(c.String("message"), c.String("key")))
	} else {
		println("Please provide a --message and a --key.")
	}
}

func HMACVerify(c *cli.Context) {
	if len(c.FlagNames()) == 3 {
		fmt.Println("verified ", VerifyHmac256(c.String("message"), c.String("tag"), c.String("key")))
	} else {
		fmt.Println("Please provide a --message with a --key and a --tag.")
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

// CURL FUNCTIONS ////////////////////////////////////////////

func SimpleCurl(c *cli.Context) {

	transport := http.Transport{Dial: dialTimeout}
	client := http.Client{Transport: &transport}

	req, _ := http.NewRequest("GET", c.String("url"), nil)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	req.Header.Set("User-Agent", "curl/7.30.0")
	req.Header.Set("Date", time.Now().UTC().Format(time.RFC1123))
	string_to_sign := req.Method + " " + req.Header["Date"][0] + " " + req.Header["Content-Type"][0]
	fmt.Println("string to sign \n" + string_to_sign)

	signed_string := ComputeHmac256(string_to_sign, "secret")
	req.Header.Add("Authorization", "Zero "+signed_string)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	fmt.Println("Request Headers")
	for key, value := range req.Header {
		fmt.Println(key, ":", value)
	}
	fmt.Println("")
	fmt.Println("Repsonse Status ", resp.Status)
	fmt.Println("")

	fmt.Println("Response Headers")
	for key, value := range resp.Header {
		fmt.Println(key, ":", value)
	}

	//	body, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Printf("%s", body)

}

var timeout = time.Duration(2 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}
