package model

import "fmt"

type Student struct {
	Name string
	age  int
	sal  float64
}

func (u *User) name() string {
	return u.Name
}
func (receiver MyNumber) name() {

}

func NewStudent(name string) *Student {
	return &Student{Name: name}
}
func (s *Student) SetAge(age int) {
	s.age = age
}

func (s *Student) GetAge() int {
	return s.age
}
func (s *Student) SetSal(f float64) {
	s.sal = f
}
func (s *Student) GetSal() float64 {
	return s.sal
}

func init() {
	fmt.Println("Student struct init-----")
}
