# General CLI helper tool playground 

> Currently playing around with HMAC in golang


###Install

`go get`


###Help

`0 -h`

```
NAME:
   zero - Help me help you!

USAGE:
   zero [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   curl   Curl.
   hmac-tag Tag a --message using a --key. Use '0 hmactag -h' to summont help.
   hmac-verify  Verify if a --tag is equal to a --message tagged using a --key. Use '0 hmacverify -h' to summon help.
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h   show help
   --version, -v  print the version
```


###Example

`0 hmac-tag --message "helloworld" --key "secret"`

`0 hmac-verify --message "helloworld" --key "secret" --tag "enwr9BlzSJvjsxitLxbHX8h1w0De7LEqP3myi7cTXJc="`
