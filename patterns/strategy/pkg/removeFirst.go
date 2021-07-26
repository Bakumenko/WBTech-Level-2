package pkg

import "fmt"

type RemoveFirst struct {
}

func (r *RemoveFirst) remove(slice []int) []int {
	fmt.Println("remove first")
	return slice[1:]
}
