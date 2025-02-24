package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	parentCode string = "Р"
	pairCode   string = "П"
	childCode  string = "Д"

	femaleCode string = "Ж"
	maleCode   string = "М"

	separatorCode string = "->"
)

var out io.Writer = os.Stdout

// Mark visited user
func (p *Person) Mark() {
	p.Marked = true
}

// If code is simple (without separator) then print user
func (p *Person) PrintOrCheckRelation(code string, desc string) {
	first_ins := strings.Index(code, separatorCode)

	if first_ins < 0 {
		fmt.Fprint(out, p.Name+" - "+desc+"\n")
	} else {
		p.CheckRelation(code[first_ins+len(separatorCode):], desc)
	}
}

// Process relationship code and print results
func (p *Person) CheckRelation(code string, desc string) {
	if p.Marked {
		return
	}
	p.Mark()
	first_ins := strings.Index(code, separatorCode)

	act_code := code

	if first_ins > 0 {
		act_code = code[:first_ins]
	}

	next := p
	if strings.Contains(act_code, parentCode) {
		next = p.Parent
	} else if strings.Contains(act_code, childCode) {
		next = p.Child
	} else if strings.Contains(act_code, pairCode) {
		next = p.Pair
	} else {
		return
	}
	for {
		if next == nil {
			break
		}

		if strings.Contains(act_code, femaleCode) {
			if strings.Contains(next.Gender, femaleCode) {
				next.PrintOrCheckRelation(code, desc)
			}

		} else if strings.Contains(act_code, maleCode) {
			if strings.Contains(next.Gender, maleCode) {
				next.PrintOrCheckRelation(code, desc)
			}
		} else {
			next.PrintOrCheckRelation(code, desc)
		}

		if strings.Contains(act_code, parentCode) {
			if strings.Contains(next.Gender, maleCode) {
				next = next.Pair
			} else {
				next = nil
			}
		} else if strings.Contains(act_code, childCode) {
			next = next.Next
		} else {
			next = nil
		}
	}
}
