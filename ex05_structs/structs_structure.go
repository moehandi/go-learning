package main

import "fmt"

type Student struct{
	Name string
	Age int
}

func main()  {
	// Case 1
	studentLiteral := Student{"ambo", 27}
	studentLiteral.showInfo()

	// Case 2
	var student Student
	student.Name = "bento"
	student.Age = 28
	fmt.Println(student.Name, student.Age)

	// Two case below point struct to Pointer
	// Case 3
	newStudent := new(Student)
	newStudent.Name = "cokro"
	newStudent.Age = 29
	newStudent.showInfo()

	// Case 4
	goStudent := &Student{"dodo", 30}
	fmt.Println(goStudent.Name, goStudent.Age)

}

func (s Student) showInfo() {
	fmt.Println(s.Name, s.Age)
}

func (s *Student) showInfos() {
	fmt.Println(s.Name, s.Age)
}