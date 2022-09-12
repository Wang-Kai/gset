# gset
Set for golang

# Get started

```go
package main

import (
	"fmt"

	"github.com/Wang-Kai/gset"
)

func main() {
	set := gset.NewStrSet().Add("hostname1", "hostname1", "hostname2")
	if !set.Has("hostname3") {
		println("The hostname3 dose not exist")
	}
	fmt.Printf("There are %d host", set.Len())
}
```
