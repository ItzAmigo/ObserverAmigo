package main

import "Observer/Observer"

func main() {
	headhunter := Observer.NewJobSite()
	Bob := Observer.NewPerson("Bob")
	headhunter.AddVacancies("Senior Golang Developer")
	headhunter.Subscribe(Bob)
	headhunter.AddVacancies("Junior Java Developer")
	headhunter.RemoveVacancies("Senior Golang Developer")

}
