package main

import (
	"fmt"
	"os"
	//"io"
	"bufio"
	"strings"
)

type user struct {
	name       string
	pair       *user
	firstChild *user
	next       *user
}

func (u *user) printAllChildren() {

	var child *user = u.firstChild
	for {
		if child == nil {
			break
		}
		fullname := strings.Split(child.name, " ")
		if strings.Contains(fullname[1], "Ж") {
			fmt.Println(fullname[0], "- дочь")
		} else {
			fmt.Println(fullname[0], "- сын")
		}

		child = child.next
	}
}

func (u *user) printAllRelatives() {
	fullname := strings.Split(u.pair.name, " ")
	if strings.Contains(fullname[1], "Ж") {
		fmt.Println(fullname[0], "- жена")
	} else {
		fmt.Println(fullname[0], "- муж")
	}
	u.printAllChildren()
}

func (u *user) getLastChildName() string {

	var child *user = u.firstChild
	var name string
	for {
		if child == nil {
			break
		}
		//fmt.Println(child.name)
		name = strings.Split(child.name, " ")[0]
		child = child.next
	}
	return name
}

func getUserIndex(arr []user, name string) int {
	if name == "" {
		return 0
	}

	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i].name, name) {
			//fmt.Println("Get user: ", arr[i].name)
			return i + 1
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	users := []user{}
	state := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			//fmt.Println("Empty string")
		} else if line[0] == '#' {
			if strings.Contains(line, "Имена") {
				//fmt.Println("Стадия 1: Имена")
				state = 1
			} else if strings.Contains(line, "женат") {
				//fmt.Println("Стадия 2: Супруги")
				state = 2
			} else if strings.Contains(line, "ребёнок") {
				//fmt.Println("Стадия 3: Дети")
				state = 3
			}
		} else {
			if state == 1 {
				u := new(user)
				u.name = line
				users = append(users, *u)
			} else if state == 2 {
				pair := strings.Split(line, " <-> ")
				i1 := getUserIndex(users, pair[0])
				i2 := getUserIndex(users, pair[1])
				if i1 > 0 && i2 > 0 {
					users[i1-1].pair = &users[i2-1]
					users[i2-1].pair = &users[i1-1]
				}
			} else if state == 3 {
				pair := strings.Split(line, " -> ")
				parent_i := getUserIndex(users, pair[0])
				newChild_i := getUserIndex(users, pair[1])
				if parent_i > 0 && newChild_i > 0 {

					if users[parent_i-1].firstChild == nil {
						users[parent_i-1].firstChild = &users[newChild_i-1]
						users[parent_i-1].pair.firstChild = &users[newChild_i-1]
						//fmt.Println("First child for ", pair[0], " is ", pair[1])
					} else {
						lastChild_name := users[parent_i-1].getLastChildName()
						lastChild_i := getUserIndex(users, lastChild_name)
						users[lastChild_i-1].next = &users[newChild_i-1]
						//fmt.Println("Next sibling for ", lastChild_name, " is ", pair[1])
					}
				}
			}
		}

		//fmt.Println(line)
	}

	//fmt.Println(users)

	inpt_scanner := bufio.NewScanner(os.Stdin)
	for inpt_scanner.Scan() {
		users[getUserIndex(users, inpt_scanner.Text())-1].printAllRelatives()
	}

	//users[getUserIndex(users, "Зоя")-1].printAllRelatives()
}
