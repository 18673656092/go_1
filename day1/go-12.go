package main

import "fmt"

type PersonInfo struct {
	Name string
	ID int
	Address string
}

func (p *PersonInfo) Running()  {

}

type IRunning interface {
	Running()
}

func main()  {
	var personDB = make(map[string]PersonInfo,100)
	personDB["Tom"] = PersonInfo{"Tom",123,"Rom322"}
	personDB["Jack"] = PersonInfo{"Jack",124,"Rom333"}
	delete(personDB,"Jack")
	person,ok :=personDB["Jack"]
	if ok {
		fmt.Println("Found person",person.Name,"with Address",person.Address)
	}else {
		fmt.Println("Not Found")
	}
}