// package main

// import "fmt"

// type Article struct {
// 	Title  string
// 	Author string
// }

// func (a Article) String() string {
// 	return fmt.Sprintf("The %q Article was written by %s", a.Title, a.Author)
// }

// type Stringer interface {
// 	String() string
// }

// func main() {
// 	s := Article{
// 		Title:  "Go Interfaces",
// 		Author: "John Doe",
// 	}
// 	Print(s.String())
// }

// func Print(s string) {
// 	fmt.Println(s)
// }
