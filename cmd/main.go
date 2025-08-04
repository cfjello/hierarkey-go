package main

import (
	"fmt"
	"os"

	"github.com/cfjello/hierarkey-go/hierarkey/pkg/hierarkey/pkg/examples"
)

func main() {
	// Parse command line arguments
	if len(os.Args) < 2 {
		// Default is example A
		fmt.Println("Running Example A...\n---------------")
		examples.RunExampleA()
		return
	}

	example := os.Args[1]

	switch example {
	case "A":
		fmt.Println("Running Example A...\n---------------")
		examples.RunExampleA()
	case "B":
		fmt.Println("Running Example B...\n---------------")
		examples.RunExampleB()
	case "C":
		fmt.Println("Running Example C...\n---------------")
		examples.RunExampleC()
	default:
		fmt.Println("Running Example A...\n---------------")
		examples.RunExampleA()
	}
}
