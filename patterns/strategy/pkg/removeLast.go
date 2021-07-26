package pkg

import "fmt"

type RemoveLast struct {
}

func (r *RemoveLast) remove(slice []int) []int {
	fmt.Println("remove last")
	return slice[:len(slice)-1]
}
