package main

import (
	"fmt"
	"slices"
	"strings"
)

func addCourses(slice []string, courses []string) []string {
	return append(slice, courses...)
}

func deleteCourses(slice []string, i int, j int) []string {
	return slices.Delete[[]string](slice, i, j)
}

func sayHello(text string, cb func(string)) {
	cb(text)
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
		courses: []map[string][]string{{"Computer Courses": {"Software Development Techniques"}}, {"Electrical Courses": {"Electrical Devices", "Electrical Machines", "Gibberish", "I love you Daddy"}}},
	}

	students.courses[0]["Computer Courses"] = addCourses(students.courses[0]["Computer Courses"], []string{"Hello World", "Lorem Ipsum", "Faithful King"})

	students.courses[1]["Electrical Courses"] = deleteCourses(students.courses[1]["Electrical Courses"], 1, 2)

	fmt.Println("courses", students.courses[0]["Computer Courses"])
	fmt.Println("courses", students.courses[1]["Electrical Courses"])

	text := "Hello World"

	sayHello(text, func(text string) {
		fmt.Println(text)
	})

	firstName := strings.Split(students.name, " ")[0]

	fmt.Println(firstName)

	switch {
	case students.level < 350:
		fmt.Println("Wida you")
	case students.level == 400:
		fmt.Println("How are you")
	}

}
