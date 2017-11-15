package alpha

import (
	"fmt"

	"github.com/matthewrudy/dep-example/beta"
)

func Call() {
	fmt.Println("Alpha: I am Alpha")
	beta.Call()
}
