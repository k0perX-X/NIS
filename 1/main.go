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
	Parent *Person
	Pair   *Person
	Child  *Person
	Next   *Person
	Marked bool
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

/*
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
	if p.Pair != nil {
		if p.Pair.Gender == "Ж" {
			fmt.Println(p.Pair.Name + " - жена")
		} else {
			fmt.Println(p.Pair.Name + " - муж")
		}
	}
	p.printAllChildren()
}*/

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
	if indexP > 0 {
		user[indexC].Parent = &user[indexP]
	}
}

type RelationDesc struct {
	Description string
	Path        string
}

var relationList []RelationDesc

func addRelation(desc string, path string) {
	r := RelationDesc{
		Description: desc,
		Path:        path,
	}
	relationList = append(relationList, r)
}

func clearMarks() {
	for index := range user {
		user[index].Marked = false
	}
}

func (p *Person) mark() {
	p.Marked = true
}

func (p *Person) printOrCheckRelation(code string, desc string) {
	first_ins := strings.Index(code, "->")

	if first_ins < 0 {
		fmt.Println(p.Name + " - " + desc)
	} else {
		p.checkRelation(code[first_ins+2:], desc)
	}
}

func (p *Person) checkRelation(code string, desc string) {
	if p.Marked {
		return
	}
	p.mark()
	first_ins := strings.Index(code, "->")

	act_code := code

	if first_ins > 0 {
		act_code = code[:first_ins]
	}

	next := p
	if strings.Contains(act_code, "Р") {
		next = p.Parent
	} else if strings.Contains(act_code, "Д") {
		next = p.Child
	} else if strings.Contains(act_code, "П") {
		next = p.Pair
	} else {
		return
	}
	for {
		if next == nil {
			break
		}

		if strings.Contains(act_code, "Ж") {
			if strings.Contains(next.Gender, "Ж") {
				next.printOrCheckRelation(code, desc)
			}

		} else if strings.Contains(act_code, "М") {
			if strings.Contains(next.Gender, "М") {
				next.printOrCheckRelation(code, desc)
			}
		} else {
			next.printOrCheckRelation(code, desc)
		}

		if strings.Contains(act_code, "Р") {
			if strings.Contains(next.Gender, "М") {
				next = next.Pair
			} else {
				next = nil
			}
		} else if strings.Contains(act_code, "Д") {
			next = next.Next
		} else {
			next = nil
		}
	}
}

func (p *Person) printAllRelationsFromDSL() {
	for index := range relationList {
		clearMarks()
		p.checkRelation(relationList[index].Path, relationList[index].Description)
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

	//addRelation("внучка", "Д->ДЖ")
	//addRelation("брат", "Р->ДМ")
	//addRelation("сестра", "Р->ДЖ")
	//addRelation("мама", "РЖ")
	//addRelation("папа", "РМ")
	//addRelation("дочь", "ДЖ")
	//addRelation("сын", "ДМ")
	//addRelation("жена", "ПЖ")
	//addRelation("муж", "ПМ")
	//addRelation("бабушка", "Р->РЖ")
	//addRelation("дедушка", "Р->РМ")
	//addRelation("внук", "Д->ДМ")
	//addRelation("теща", "ПЖ->РЖ")
	//addRelation("свекр", "ПМ->РМ")
	//addRelation("шурин", "ПЖ->Р->ДМ")

	//_, err := Parse("input.txt", []byte(input))
	_, err := ParseFile("input.txt")
	if err != nil {
		log.Fatalf("Error parsing input: %s", err)
	}

	_, err = ParseFile("relationships.txt")
	if err != nil {
		log.Fatalf("Error parsing input: %s", err)
	}

	for index := range relationList {
		fmt.Println(relationList[index].Description, relationList[index].Path)
	}

	inptScanner := bufio.NewScanner(os.Stdin)
	for inptScanner.Scan() {
		//user[getUserIndex(user, inptScanner.Text())].printAllRelatives()
		user[getUserIndex(user, inptScanner.Text())].printAllRelationsFromDSL()
	}

}
