package main

import (
	"fmt"
	"io"
	"os"
    "strings"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "crypto/subtle"
)

//usage is a helpful usage string shown when
//the user doesn't supply enough arguments.
//The `...` allows line breaks within the string
const usage = `
USAGE:
	hmac sign signing-key < file-to-sign
	hmac verify signing-key signature < file-to-verify
`

//exit codes
//iota resets to 0 every time you start
//a const block, and increments by 1
//each time you use it
const (
	_                        = iota //ignore first usage (which is zero)
	exitCodeUsage                   // = 1
	exitCodeProcessing              // = 2
	exitCodeInvalidSignature        // = 3
)

//showUsage shows the usage string and exits
//with the code exitCodeUsage
func showUsage() {
	fmt.Println(usage)
	os.Exit(exitCodeUsage)
}

//sign returns a base64-encoded HMAC signature given a
//signingKey and a read stream. It returns an error if
//there was an error reading from the stream.
func sign(signingKey string, stream io.Reader) (string, error) {
    key := []byte(signingKey)

    //create a new HMAC hasher
    h := hmac.New(sha256.New, key)

    if _, err := io.Copy(h, stream); err != nil {
        return "", fmt.Errorf("error copying bytes: %v", err)
    }
    sig := h.Sum(nil)

    //TODO: implement this function according to the comments
	//HINTS:
	//https://drstearns.github.io/tutorials/sessions/#secdigitalsignatureswithhmac
	//https://golang.org/pkg/crypto/hmac/
	//https://golang.org/pkg/io/#Copy
	//https://golang.org/pkg/encoding/base64/
	return base64.URLEncoding.EncodeToString(sig), nil
}

//verify returns true if the base64-encoded HMAC `signature`
//matches the bytes read from `stream`, or false if otherwise.
//If there is an error decoding the base64 signature, or reading
//from `stream`, this will return false and the error.
func verify(signingKey string, signature string, stream io.Reader) (bool, error) {
	//TODO: implement this function according to the comments
	//HINTS: (same as above plus the following)
	//https://golang.org/pkg/crypto/subtle/#ConstantTimeCompare
    key := []byte(signingKey)

    sigBytes, err := base64.URLEncoding.DecodeString(signature)
    if err != nil {
        return false, fmt.Errorf("decoding string failed: %v", err)
    }
    //create a new HMAC hasher
    h := hmac.New(sha256.New, sigBytes)
    if _, err := io.Copy(h, stream); err != nil {
        return false, fmt.Errorf("copying bytes failed: %v", err)
    }

    newSig := h.Sum(nil)
	return subtle.ConstantTimeCompare(newSig, sigBytes) == 1, nil
}

func main() {
	if len(os.Args) < 3 {
		showUsage()
	}

	//TODO: use the os.Args slice
	//to get the command and signing key
	//then switch on the command
	//and call the appropriate function above,
	//printing the return value or error
	command := strings.ToLower(os.Args[1])
	signingKey := os.Args[2]


	switch command {
    case "sign": {
        sig, err := sign(signingKey, os.Stdin)
        if err != nil {
            fmt.Printf("Error signing: %v", err)
            os.Exit(exitCodeProcessing)
        }

        fmt.Println(sig)
    }
    case "verify": {
        sig64 := os.Args[3]
        if len(sig64) == 0 {
            showUsage()
        }

        valid, err := verify(signingKey, sig64, os.Stdin)
        if err != nil {
            fmt.Printf("Error verifying: %v", err)
            os.Exit(exitCodeProcessing)
        }

        if valid {
            fmt.Println("Signatutre was valid!")
        } else {
            fmt.Println("invalid sig")
        }
    }
    default: showUsage()
    }
}
