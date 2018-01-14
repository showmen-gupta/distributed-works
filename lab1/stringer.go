// +build !solution

package lab1

import "strconv"
import "strings"

/*
Task 2: Stringers

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}

A Stringer is a type that can describe itself as a string. The fmt package (and
many others) look for this interface to print values.

Implement the String() method for the Student struct.

A struct

Student{ID: 42, FirstName: John, LastName: Doe, Age: 25}

should be printed as

"Student ID: 42. Name: Doe, John. Age: 25.
*/

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func (s Student) String() string {
	reg := []string{s.LastName, s.FirstName}
	var id = strconv.Itoa(s.ID)
	var name = strings.Join(reg[:], ", ") //joining strings of name
	var age = strconv.Itoa(s.Age)

	studentreg := []string{"Student ID:" + " " + id, " Name:" + " " + name, " Age:" + " " + age + "."}
	var output = strings.Join(studentreg[:], ".") //final joining of all the strings
	return output
}
