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

type RelationDesc struct {
	Description string
	Path        string
}

var inputFileName = "input.txt"
var relationshipFileName = "relationships.txt"

var head *Person = nil
var tail *Person = nil

var user []Person
var relationList []RelationDesc

func init() {
	user = make([]Person, 0)
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
		if user[indexP].Pair != nil {
			user[indexP].Pair.Child = &user[indexC]
		}
	} else {
		lastKidName_idx := getUserIndex(user, user[indexP].GetLastChildName())
		user[lastKidName_idx].Next = &user[indexC]
	}
	if indexP > 0 {
		user[indexC].Parent = &user[indexP]
	}
}

func addRelation(desc string, path string) {
	r := RelationDesc{
		Description: desc,
		Path:        path,
	}
	relationList = append(relationList, r)
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

func clearMarks() {
	for index := range user {
		user[index].Marked = false
	}
}

func (p *Person) GetLastChildName() string {

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

func (p *Person) PrintAllRelationsFromList(relationships []RelationDesc) {
	for index := range relationships {
		clearMarks()
		p.CheckRelation(relationships[index].Path, relationships[index].Description)
	}
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

	_, err := ParseFile(inputFileName)
	if err != nil {
		log.Fatalf("Error parsing input file: %s", err)
	}

	_, err = ParseFile(relationshipFileName)
	if err != nil {
		log.Fatalf("Error parsing relationship file: %s", err)
	}

	for index := range relationList {
		fmt.Println(relationList[index].Description, relationList[index].Path)
	}

	inptScanner := bufio.NewScanner(os.Stdin)
	for inptScanner.Scan() {
		user[getUserIndex(user, inptScanner.Text())].PrintAllRelationsFromList(relationList)
	}

}
