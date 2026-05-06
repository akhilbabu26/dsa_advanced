package main

import "fmt"

type Set map[int]bool

func (s Set) Add(value int){
	s[value] = true
}

func (s Set) Remove(value int){
	delete(s, value)
}

func (s Set) Contains(value int) bool{
	return s[value]
}

func main(){
	s := make(Set)

	s.Add(20)
	s.Add(40)
	s.Add(20)
	s.Add(60)

	fmt.Println(s.Contains(20))
	fmt.Println(s.Contains(60))
	fmt.Println(s.Contains(20))
	fmt.Println(s.Contains(2))

	s.Remove(20)
	fmt.Println(s.Contains(20))
}

