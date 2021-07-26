package pkg

import "fmt"

type Reception struct {
	next department
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
	} else {
		p.registrationDone = true
		fmt.Println("Reception registering patient")
	}
	r.next.Execute(p)
}

func (r *Reception) SetNext(next department) {
	r.next = next
}
