package main

import (
	"fmt"
)

// type Doctor struct {
// 	number     int // lowercase, internal to the package
// 	actorName  string
// 	companions []string
// }

// type Animal struct { //TAG
// 	Name   string `required max:"100"`
// 	Origin string
// }
// type Bird struct {
// 	Animal   // Animal struct is embedded here, similar to extending
// 	SpeedKPH float32
// 	CanFly   bool
// }

// var ( // for arranging some variables
// 	actorName    string = "Sladen"
// 	companion    string = "sarah"
// 	doctorNumber int    = 3
// 	seasonNumber int    = 11
// )
// var (
// 	counter int = 0
// )

// const (
// 	a = iota //0
// 	b = iota //1
// 	c = iota //2
// )
// const (
// 	a2 = iota //0
// )

// var i int = 27 // package level visibility
// var I int = 27 // globally visible, capital letters declared

func main() {

	// fmt.Printf("%V\n", a)
	// fmt.Printf("%V\n", b)
	// fmt.Printf("%V\n", c)
	// fmt.Printf("%V\n", a2)

	// var i int // block level visility
	// i = 42
	// var j int = 45
	// var l float32 = 45.2552
	// k := 99
	// m := 99 // variale declared so it must be used in go unless results in error
	// fmt.Println("Hello, Playground")
	// fmt.Println(i, j, k, m)
	// fmt.Printf("%V, %T", l, l) // value, type

	// //var theURL string = "https://google.com"
	// //var HttpRequest string = ""

	// var n int = 42
	// fmt.Printf("%V, %T\n", n, n)

	// var o float32
	// o = float32(n)
	// fmt.Printf("%V, %T\n", o, o)
	// var p string
	// p = strconv.Itoa(n)
	// fmt.Printf("%V, %T\n", p, p)

	// // var l bool = false
	// // n := 1 == 1
	// // m := 1 == 2

	// // fmt.Printf("%V, %T\n", n, l)
	// // fmt.Printf("%V, %T\n", m, m)

	// // var o uint16 = 42
	// // fmt.Printf("%V, %T\n", o, o)

	// // var c complex64 = 1 + 2i
	// // fmt.Printf("%V, %T\n", c, c)
	// // fmt.Printf("%V, %T\n", real(c), real(c))
	// // fmt.Printf("%V, %T\n", imag(c), imag(c))

	// // s := "this is a string"
	// // b := []byte(s)
	// // fmt.Printf("%V, %T\n", b, b)

	// // r := 'a' // rune in go declared with single quotes
	// // fmt.Printf("%V, %T\n", r, r)

	// grades := [3]int{97, 85, 93}

	// grades2 := [...]int{97, 85, 93}
	// var students [3]string
	// students[0] = "Lisa"
	// fmt.Printf("%V\n", grades)
	// fmt.Printf("%V\n", grades2)
	// fmt.Printf("%V\n", students)

	// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// fmt.Println(identityMatrix)

	// a := [...]int{1, 2, 3}
	// b := a
	// //b := &a
	// b[1] = 5
	// fmt.Println(a, len(a), cap(a))
	// fmt.Println(b)

	// /// map
	// statePopulations := map[string]int{
	// 	"California": 39250017,
	// 	"Texas":      27862596,
	// 	"Ohio":       11614373,
	// }
	// statePopulations["Georgia"] = 10310371
	// delete(statePopulations, "Texas")
	// fmt.Println(statePopulations, len(statePopulations))
	// pop, ok := statePopulations["Ohioj"]
	// fmt.Println(pop, ok)

	// //structs
	// aDoctor := Doctor{
	// 	number:    3,
	// 	actorName: "Jon",
	// 	companions: []string{
	// 		"name1",
	// 		"name2",
	// 		"name3",
	// 	},
	// }
	// fmt.Println(aDoctor)
	// fmt.Println(aDoctor.actorName)

	// newDoctor := struct{ name string }{name: "John Pertwee"}
	// //anotherDoctor := newDoctor
	// anotherDoctor := &newDoctor
	// anotherDoctor.name = "Tom Baker"
	// fmt.Println(newDoctor, anotherDoctor)

	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// fmt.Println(b, b.Name)
	// newBird := Bird{
	// 	Animal:   Animal{Name: "new bird", Origin: "USA"},
	// 	SpeedKPH: 48,
	// 	CanFly:   false,
	// }
	// fmt.Println(newBird)

	// //TAG
	// tag := reflect.TypeOf(Animal{})
	// field, _ := tag.FieldByName("Name")
	// fmt.Println(field.Tag) //required max:"100"

	var s []int
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = a[2:4]
	newS := append(s, 55, 66)
	fmt.Print(len(newS), cap(newS))
}
