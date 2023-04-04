package shortener

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

	"github.com/itchyny/base58-go"
)

var spf = fmt.Sprintf

func errChecker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sha2560f(input string) []byte {

	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {

	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)

	errChecker(err)

	return string(encoded)
}

func GenerateShortLink(initialLink string, userID string) string {

	// Hashing  initialUrl + userId url with sha256.
	urlHashBytes := sha2560f(initialLink + userID)

	// Derive a big integer number from the hash bytes generated during the hasing.
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// Finally apply base58  on the derived big integer value
	finalString := base58Encoded([]byte(spf("%d", generatedNumber)))

	// and pick the first 8 characters
	return finalString[:8]
}
