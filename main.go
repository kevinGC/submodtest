package submodtest

import (
	"fmt"

	"github.com/kevinGC/submodtest/dep"
)

// func main() {
// 	fmt.Printf("Hello, %s\n", dep.Message())
// }

func Message() string {
	return fmt.Sprintf("Hello, %s\n", dep.Message())
}
