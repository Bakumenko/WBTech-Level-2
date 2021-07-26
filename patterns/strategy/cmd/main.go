package main

import "strategy/pkg"

func main() {
	oleg := pkg.NewPersonSlice("Oleg")
	sasha := pkg.NewPersonSlice("Sasha")
	oleg.SetSlice([]int{1, 2, 3})
	sasha.SetSlice([]int{1, 2, 3})

	removeFirst := &pkg.RemoveFirst{}
	removeLast := &pkg.RemoveLast{}
	oleg.SetRemoveAlgorithm(removeFirst)
	sasha.SetRemoveAlgorithm(removeLast)

	oleg.Remove()
	sasha.Remove()
}
