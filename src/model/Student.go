package model

type student struct {
	Name string
	age  int
	sal  float64
}

func NewStudent(name string) *student {
	return &student{Name: name}
}
func (s *student) SetAge(age int) {
	s.age = age
}

func (s *student) GetAge() int {
	return s.age
}
func (s *student) SetSal(f float64) {
	s.sal = f
}
func (s *student) GetSal() float64 {
	return s.sal
}
