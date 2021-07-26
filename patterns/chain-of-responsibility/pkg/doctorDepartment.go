package pkg

import "fmt"

type Doctor struct {
	next department
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
	} else {
		p.doctorCheckUpDone = true
		fmt.Println("Doctor checking patient")
	}
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next department) {
	d.next = next
}
