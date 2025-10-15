/*
main cli entry
*/
package main

import (
	"fmt"

	"github.com/siakhooi/picsum/internal/version"
)

func main() {
	fmt.Printf("picsum %s\n", version.GetVersion())
}
