
# hkCrypto

## Usage

```go
package main

import (
	"fmt"
	"go.hikit.io/hkcrypto"
)

func main() {
	code := hkcrypto.NewAesCbcEncode([32]byte{}, [16]byte{}, hkcrypto.Strength256)
	edata := code.Encrypt([]byte("hikit.io"))
	fmt.Println(string(edata))
	ddata := code.Decrypt([]byte(edata))
	fmt.Println(string(ddata))
}
```