package pkg

import "fmt"

type Nurse struct {
	next department
}

func (d *Nurse) Execute(p *Patient) {
	if p.nurseGaveCertificate {
		fmt.Println("The nurse has already issued a certificate")
	} else {
		p.nurseGaveCertificate = true
		fmt.Println("The nurse gave a certificate")
	}
}

func (d *Nurse) SetNext(next department) {
	d.next = next
}
