package pkg

import "fmt"

type PersonSlice struct {
	person          string
	slice           []int
	removeAlgorithm iRemove
}

func NewPersonSlice(name string) *PersonSlice {
	return &PersonSlice{person: name}
}

func (p *PersonSlice) SetSlice(s []int) {
	p.slice = s
	fmt.Println(p.person, p.slice)
}

func (p *PersonSlice) SetRemoveAlgorithm(algo iRemove) {
	p.removeAlgorithm = algo
}

func (p *PersonSlice) Remove() {
	if p.removeAlgorithm != nil {
		p.slice = p.removeAlgorithm.remove(p.slice)
		fmt.Println(p.person, p.slice)
	} else {
		fmt.Println("set algorithm of remove")
	}
}
