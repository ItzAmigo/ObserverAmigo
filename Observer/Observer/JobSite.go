package Observer

import _ "fmt"

type JobSite struct {
	observerList []Observer
	vacancyList  []string
}

func NewJobSite() *JobSite {
	j := new(JobSite)
	j.observerList = make([]Observer, 0, 10)
	j.vacancyList = make([]string, 0, 10)
	return j
}

func (j *JobSite) AddVacancies(vacancy string) *JobSite {
	j.vacancyList = append(j.vacancyList, vacancy)
	return j
}

func remove[T any](slice []T, i int) []T {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func (j *JobSite) RemoveVacancies(vacancy string) {
	for i := range j.vacancyList {
		if j.vacancyList[i] == vacancy {
			j.vacancyList = remove(j.vacancyList, i)
			break
		}
	}
	j.sendAll()
}

func (j *JobSite) Subscribe(obs Observer) {
	j.observerList = append(j.observerList, obs)
}

func (j *JobSite) Unsubscribe(obs Observer) {
	for i := range j.observerList {
		if j.observerList[i] == obs {
			j.observerList = remove(j.observerList, i)
			break
		}
	}
}

func (j *JobSite) sendAll() {
	for _, observer := range j.observerList {
		observer.handleEvent(j.vacancyList)
	}
}
