# Custom signals with handlers


```go
package main

import (
	"fmt"
    "github.com/onrik/sigi"
)

func CountHandler(count int) {
	fmt.Println("New count:", count)
}

func main() {
	sigi.Connect("count-changed", CountHandler)

	sigi.Emit("count-changed", 1)
	sigi.Emit("count-changed", 2)

	sigi.Disconnect("count-changed", CountHandler)
	
	// Now to disconnect anonymous handler
	connector, err := sigi.Connect("count-changed", func(count int) {
	    // ...
	})
	
	connector.Disconnect()
}


```
