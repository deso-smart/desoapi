# DeSo client

## TODO

- [ ] Ability to upload files
- [ ] Refactoring `executeRequest` method in `Client`
- [ ] Tests

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	desoRoutes "github.com/deso-smart/deso-backend/v3/routes"
	"github.com/deso-smart/desoapi"
	"log"
)

func main() {
	client, err := desoapi.NewClient("https://node.deso.org")
	if err != nil {
		log.Fatal(err)
	}
	payload := &desoRoutes.GetSingleProfileRequest{
		PublicKeyBase58Check: "...",
	}
	resp, err := client.GetSingleProfile(payload)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))
}
```