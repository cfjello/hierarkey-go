package examples

import (
	"fmt"

	"github.com/cfjello/hierarkey-go/hierarkey/pkg/hierarkey/pkg/hierarkey"
)

func RunExampleC() {
	hk := hierarkey.NewHierarKey(1, 4, "")

	fmt.Println("Get the root leaf:")
	fmt.Println(hk.NextLeaf())
	fmt.Println("Go up a few levels:")
	fmt.Println(hk.NextLevel())
	fmt.Println(hk.NextLevel())
	fmt.Println(hk.NextLeaf())
	fmt.Println(hk.NextLevel())
	fmt.Println("Go down a few levels:")
	fmt.Println(hk.PrevLevel())
	fmt.Println(hk.PrevLevel())
	fmt.Println("Jump to an existing level:")
	fmt.Println(hk.JumpToLevel("0001.0001.0002"))
	fmt.Println("Jump to an arbitrary level:")
	fmt.Println(hk.JumpToLevel("7.6.5"))
	fmt.Println(hk.NextLeaf())
	fmt.Println("Go a down 2 levels:")
	fmt.Println(hk.PrevLevel(2))
	fmt.Println("Jump to a level in between:")
	fmt.Println(hk.JumpToLevel("2.1"))
	fmt.Println(hk.NextLeaf())
}
