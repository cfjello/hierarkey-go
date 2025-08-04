package main

import (
	"fmt"
	"os"

	"github.com/cfjello/hierarkey-go/pkg/examples"
)

func main() {
	// Parse command line arguments
	example := "A" // Default example
	if len(os.Args) < 2 {
		// Default is example A
		fmt.Println("Running Example A...\n---------------")
		examples.RunExampleA()
		return
	} else {
		example = os.Args[1]
	}

	if len(example) > 1 {
		fmt.Println("Invalid example argument. Use 'A', 'B', or 'C'. Defaulting to Example A.")
		example = "A"
	}

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
