package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Name   string
	Gender string
	Pair   *Person
	Child  *Person
	Next   *Person
}

var head *Person = nil
var tail *Person = nil

var user []Person

func init() {
	user = make([]Person, 0)
}

func (p *Person) getLastChildName() string {

	var child *Person = p.Child
	var name string
	for {
		if child == nil {
			break
		}
		name = child.Name
		child = child.Next
	}
	return name
}

func (p *Person) printAllChildren() {

	var child *Person = p.Child
	for {
		if child == nil {
			break
		}
		if child.Gender == "Ж" {
			fmt.Println(child.Name + " - дочь")
		} else {
			fmt.Println(child.Name + " - сын")
		}

		child = child.Next
	}
}

func (p *Person) printAllRelatives() {
	if p.Pair.Gender == "Ж" {
		fmt.Println(p.Pair.Name + " - жена")
	} else {
		fmt.Println(p.Pair.Name + " - муж")
	}
	p.printAllChildren()
}

func addPerson(name, gender string) {
	p := &Person{
		Name:   name,
		Gender: gender,
	}
	user = append(user, *p)
}

func addPair(pairNameA, pairNameB string) {
	indexA := getUserIndex(user, pairNameA)
	indexB := getUserIndex(user, pairNameB)

	user[indexA].Pair = &user[indexB]
	user[indexB].Pair = &user[indexA]
}

func addChildren(parentName, childName string) {
	indexP := getUserIndex(user, parentName)
	indexC := getUserIndex(user, childName)

	if user[indexP].Child == nil {
		user[indexP].Child = &user[indexC]
		user[indexP].Pair.Child = &user[indexC]
	} else {
		lastKidName_idx := getUserIndex(user, user[indexP].getLastChildName())
		user[lastKidName_idx].Next = &user[indexC]
	}

}

func getUserIndex(arr []Person, name string) int {
	if name == "" {
		return 0
	}

	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i].Name, name) {
			return i
		}
	}
	return 0
}

func main() {

	//_, err := Parse("input.txt", []byte(input))
	_, err := ParseFile("input.txt")
	if err != nil {
		log.Fatalf("Error parsing input: %s", err)
	}

	inptScanner := bufio.NewScanner(os.Stdin)
	for inptScanner.Scan() {
		user[getUserIndex(user, inptScanner.Text())].printAllRelatives()
	}

}
