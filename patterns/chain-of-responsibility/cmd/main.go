package main

import "chain-of-responsibility/pkg"

func main() {
	nurse := &pkg.Nurse{}

	doctor := &pkg.Doctor{}
	doctor.SetNext(nurse)

	reception := &pkg.Reception{}
	reception.SetNext(doctor)

	patient := &pkg.Patient{Name: "Oleg"}

	reception.Execute(patient)
}
