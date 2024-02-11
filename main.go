package main

import (
	"fmt"
	"strings"
)

func addCourses(slice []string, courses []string) []string {
	return append(slice, courses...)
}

func main() {
	type Students struct {
		name    string
		dept    string
		faculty string
		level   uint
		courses []map[string][]string
	}

	students := Students{
		name:    "Faithful Eromosele",
		dept:    "Computer",
		faculty: "Engineering",
		level:   300,
		courses: []map[string][]string{{"Computer Courses": {"Software Development Techniques"}}, {"Electrical Courses": {"Electrical Devices", "Electrical Machines"}}},
	}

	students.courses[0]["Computer Courses"] = addCourses(students.courses[0]["Computer Courses"], []string{"Hello World", "Lorem Ipsum", "Faithful King"})

	fmt.Println("courses", students.courses[0]["Computer Courses"])

	firstName := strings.Split(students.name, " ")[0]

	fmt.Println(firstName)

	switch {
	case students.level < 350:
		fmt.Println("Wida you")
	case students.level == 400:
		fmt.Println("How are you")
	}

}
