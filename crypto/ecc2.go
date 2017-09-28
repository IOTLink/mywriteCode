package main
import (
	"crypto/x509"
	"os"
	"fmt"
	"encoding/pem"
)

func main() {
var pemBytes = []byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIKGOgzn9u8RCSwwJj0sGOog6QGpDNkCuBRNsv76bRXLYoAoGCCqGSM49
AwEHoUQDQgAEPAYLQF6I4NQ1Q0AjeHqJj7fDX/WwJ6xba5aDQ7V9pIQfq8k+JUME
RUBF85MS+jPu5Rn+59AP9aPRSybIQsxZrg==
-----END EC PRIVATE KEY-----
`)

	block, rest := pem.Decode(pemBytes)
	if rest != nil {
		fmt.Println(rest)
	}

	privatekey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Public Key :  : %x\n\n", privatekey.PublicKey)
	fmt.Printf("Private Key D :  : %x\n\n", privatekey.D)
}