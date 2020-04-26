package main

import (
	"fmt"

	"github.com/ob-algdatii-20ss/ob-algdatii-20ss/misc/unionfind"
)

func main() {
	k := unionfind.NewCollection()

	for i := 1; i <= 6; i++ {
		k.MakeSet(i)
	}
	k.Union(1, 2)
	k.Union(3, 4)
	k.Union(3, 5)
	k.Union(1, 3)

	for i := 1; i <= 6; i++ {
		fmt.Printf("%d -> %d (Size: %d)\n", i, k[i].Parent, k[i].Size)
	}

	for i := 1; i <= 6; i++ {
		fmt.Printf("Find(%d) = %d\n", i, k.Find(i))
	}
}
