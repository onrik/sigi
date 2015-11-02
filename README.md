# Custom signals with handlers


```go
package main

import (
    "github.com/onrik/sigi"
	"net/http"
)

func CountHandler(count int) {
	fmt.Println("New count:", count)
}

func main() {
	sigi.Connect("count-changed", CountHandler)

	sigi.Emit("count-changed", 1)
	sigi.Emit("count-changed", 2)

	sigi.Disconnect("count-changed", CountHandler)
}


```
