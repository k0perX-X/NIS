package main

import (
	"bytes"
	//"fmt"
	"os"
	"strings"
	"testing"
)

func TestCheckDaughter(t *testing.T) {
	usr := &Person{
		Name:   "Person",
		Gender: femaleCode,
	}

	daughter := &Person{
		Name:   "Daughter",
		Gender: femaleCode,
	}

	usr.Child = daughter

	buf := &bytes.Buffer{}
	out = buf

	usr.CheckRelation("ДЖ", "Дочь")

	out = os.Stdout

	//fmt.Fprint(out, buf.String()+"\n")

	if !strings.Contains(buf.String(), "Daughter") {
		t.Errorf("Result was incorrect")
	}
}

func TestCheckGrandmother(t *testing.T) {
	usr := &Person{
		Name:   "Person",
		Gender: maleCode,
	}

	gm := &Person{
		Name:   "Grandmother",
		Gender: femaleCode,
	}

	f := &Person{
		Name:   "Father",
		Gender: maleCode,
	}

	addPerson(usr.Name, usr.Gender)
	addPerson(f.Name, f.Gender)
	addPerson(gm.Name, gm.Gender)

	addChildren(f.Name, usr.Name)
	addChildren(gm.Name, f.Name)

	buf := &bytes.Buffer{}
	out = buf

	clearMarks()
	user[getUserIndex(user, "Person")].
		CheckRelation(parentCode+separatorCode+parentCode+femaleCode, "")

	out = os.Stdout

	//fmt.Fprint(out, buf.String()+"\n")

	if !strings.Contains(buf.String(), "Grandmother") {
		t.Errorf("Result was incorrect")
	}
}
