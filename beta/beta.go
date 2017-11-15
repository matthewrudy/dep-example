package beta

import (
	"fmt"

	"github.com/matthewrudy/dep-example/gamma"
)

func Call() {
	fmt.Println("Beta:  I am Beta")
	gamma.Call()
}
