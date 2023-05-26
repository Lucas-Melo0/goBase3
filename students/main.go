package main

import (
	"fmt"
	"time"
)

type student struct {
	Name          string
	Surname       string
	Id            string
	admissionDate time.Time
}

func main() {

	lucas := student{"Lucas", "Melo", "02023487216", time.Now()}
	fmt.Println(lucas)
	lucasDetails := lucas.Details()
	println(lucasDetails)

}

func (s student) Details() string {
	return "There are no details for this student"
}
