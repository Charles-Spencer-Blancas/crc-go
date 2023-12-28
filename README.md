
# CRC-Go

This is an implementation of Cyclic Redundancy Check in Go.

# How to Run

Go into the folder and do `go run . [args]` or do `go build` which creates an executable called `crc-go`
and then `./crc-go [args]`

# Usage

Without any flags, the app generates the message to send
given a generator and data. The first argument is the generator, the second is the data.
So `./crc-go 23 4555` encodes 4555 with the generator 23

With the `-c` flag, the app checks if a message is divisible by the generator. In other words,
it checks if any detectable errors are present. `./crc-go -c 32 4576` checks if the message 4576
is divisible by the generator 32.

The `-d` flag takes a generator and a message and decodes the message (only if 
no errors are detected, of course). `./crc-go -d 7 112` decodes the message 112 encoded
with the generator 7.

The `-b` flag allows the user to input their numbers in binary instead of decimal, and 
the output of the app will also be in binary. It can be used in any of the above modes.
`./crc-go -b 111 11100` is equivalent to `./crc-go 7 28`
