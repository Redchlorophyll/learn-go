package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Book struct {
	BookName  string
	TotalPage int16
	Publisher string
}
type AuthorStruct struct {
	Name  string
	Age   int8
	Books []Book
}

type BookInterface interface {
	GetBookName() string
}

func (book *Book) GetBookName() string {
	return book.BookName
}

func getBook(param BookInterface) string {
	return "the book name is " + param.GetBookName()
}

func (book *Book) SayHello() {
	fmt.Println("hello", book.BookName)
}

func testing(hello string, name string) (string, string) {
	return hello + " " + name, name
}

func NamedReturn() (FirstName string, LastName string) {
	FirstName = "Dhonni"
	LastName = "Ari"

	return FirstName, LastName
}

func variadicFunc(nums ...int) int {
	var result = 0
	for _, num := range nums {
		result += num
	}
	return result
}

func filter(word string) string {
	if word == "anjing" {
		return "..."
	}
	return word
}

type FilterType func(string) string

func HelloName(name string, filter FilterType) string {
	return "Hello " + filter(name)
}

func FactorialRecursion(num int) int {
	if num == 1 {
		return 1
	}
	return num * FactorialRecursion(num-1)
}

func AppLog() {
	fmt.Println("app selesai dijalankan!")
	message := recover()
	if message != nil {
		fmt.Println(message)
	}

}

func StartApp(IsError bool) {
	defer AppLog()
	if IsError {
		panic("App Error")
	}
	fmt.Println("App berjalan normal")
}

func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "kosong"
	}
}

func changeBookName(book *Book) {
	book.BookName = "hello"
}

func main() {
	name := "Dhonni"
	var hello = "hello"
	fmt.Println(hello)
	name = "agustrio"
	var (
		bulk1 = "eko"
		bulk2 = "agus"
	)
	const test = "test"
	const (
		first = "agus"
		last  = "yuli"
	)
	fmt.Println(first, last)
	fmt.Println(test)
	fmt.Println(name, bulk1, bulk2)

	// convertion int
	const num32 int32 = 100000
	const num64 int64 = int64(num32)
	// const num8 int8 = int8(num32) out of int8 range
	fmt.Println(num32, num64)

	// convertion string byte to str
	const name2 = "eko kurniawan"
	const stringnum = "111"
	fmt.Println(name2, name2[0], string(name2[0]))

	i, _ := strconv.Atoi(stringnum)
	f, _ := strconv.ParseInt(stringnum, 10, 8) //string, base int, bitsize
	fmt.Println(stringnum, i, f)

	if true {
		var name string = "irian"
		fmt.Println(name)
	}

	// type declaration
	type NoKTP string
	const UserKTP NoKTP = "123123"
	fmt.Println(UserKTP)

	// array
	const ArrayLength = 3
	var FirstArray [ArrayLength]string
	var SecondArray = [2]int8{1, 2}

	FirstArray[0] = "kennedy"
	// FirstArray[1] = "Adias"
	FirstArray[2] = "gecko"

	fmt.Println(FirstArray[0])
	fmt.Println(SecondArray)
	fmt.Println(len(FirstArray)) //check array length of total available slot

	// slice
	DateArray := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"Mei",
		"Juni",
		"Juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	DateSlice1 := DateArray[4:7]
	fmt.Println(DateSlice1)
	fmt.Println("len, cap =>", len(DateSlice1), cap(DateSlice1))

	// changes in array will affect slice too
	DateArray[5] = "Juni - array updated"
	fmt.Println(DateSlice1)

	// changes in slices will affect array
	DateSlice1[0] = "Mei - Slice Update"
	fmt.Println(DateArray)

	// append method will affect array if capacity still sufficient
	// otherwise it will create new array and previous array will not affected to changes in new slice

	// e.g 1
	DateSlice2 := append(DateSlice1, "hello")
	fmt.Println("=========eg 1 append==============")
	fmt.Println(DateSlice2)
	fmt.Println(DateArray) // will replace august to hello

	// e.g 2
	DateSlice3 := append(DateArray[11:], "ramsay", "julius")
	fmt.Println("========eg 2 append===============")
	fmt.Println(DateSlice3, len(DateSlice3), cap(DateSlice3))
	fmt.Println(DateArray)

	// create new slice
	NewSlice := make([]string, 2, 5)
	fmt.Println("========make new slice===============")
	NewSlice[0] = "eko"
	NewSlice[1] = "Julius"

	//  to add value to array in slice use append
	NewSlice = append(NewSlice, "Jerid")
	fmt.Println(NewSlice)

	// copy slice
	fmt.Println("========make new slice===============")
	NewSlice2 := make([]string, len(NewSlice), cap(NewSlice))
	copy(NewSlice2, NewSlice)
	NewSlice2 = append(NewSlice2, "gill")
	// copy wont update copied slice
	fmt.Println(NewSlice2)
	fmt.Println(NewSlice)

	// fastest way to create slice

	// this is array
	ThisIsArr := [2]string{"Dhonni", "Ari"}
	// this is slice, no definition for length so it is slice
	ThisIsSlice := []string{"julius", "Kounde"}

	fmt.Println("========make new slice - faster and comparation to creating arr===============")
	fmt.Println(ThisIsArr, ThisIsSlice)

	// map in go
	// simple map
	SimpleMap := map[string]string{
		"name": "dhonni",
		"age":  "40",
	}

	fmt.Println(SimpleMap)

	BookAuthor := AuthorStruct{
		Name: "Julia",
		Age:  20,
		Books: []Book{
			{
				BookName:  "Everlasting sun",
				TotalPage: 20,
				Publisher: "august",
			},
		},
	}
	fmt.Println("=================map===========")
	fmt.Println(BookAuthor.Books[0])

	// for
	fmt.Println("=================for===========")
	arr := [3]string{"dhonni", "ari", "hendra"}
	for i := 0; i < len(arr); i += 1 {
		fmt.Println(arr[i])
	}

	for idx, name := range arr {
		fmt.Println(idx+1, name)
	}

	word, name := testing("hello", "dhonni")
	fmt.Println(word, name)

	// named return
	a, b := NamedReturn()
	fmt.Println("=================named return===========")
	fmt.Println(a, b)

	// variadic
	fmt.Println("=================variadic func===========")
	Output1 := variadicFunc(1, 2, 3, 4, 5)
	fmt.Println(Output1)
	VariadicSlice := []int{1, 2, 3, 4, 5, 6, 7}
	Output2 := variadicFunc(VariadicSlice...)
	fmt.Println(Output2)

	// function as variable
	fmt.Println("=================func as var===========")
	NewVariadicFunc := variadicFunc
	fmt.Println(NewVariadicFunc(5, 5))

	// function as parameter
	fmt.Println("=================func as params===========")
	fmt.Println(HelloName("Dhonni", filter))
	fmt.Println(HelloName("anjing", filter))

	fmt.Println("=================anon func===========")
	AnonFunc := func() string {
		return "hello"
	}

	fmt.Println(AnonFunc())

	// recursion
	fmt.Println("=================recursion===========")
	fmt.Println(FactorialRecursion(3))

	// closure
	fmt.Println("=================closure===========")
	konter := 0

	Increment := func() {
		konter++
	}

	Increment()
	Increment()
	fmt.Println(konter)

	// defer, panic, recover
	fmt.Println("=================defer, panic, recover===========")
	StartApp(false)
	fmt.Println("......")
	StartApp(true)

	fmt.Println("=================construct method===========")
	rully := Book{
		BookName:  "clean architecture",
		TotalPage: 300,
		Publisher: "siaga medika",
	}

	rully.SayHello()

	fmt.Println("=================interface===========")
	//   adi := Book{
	// 		BookName:  "clean architecture",
	// 		TotalPage: 300,
	// 		Publisher: "siaga medika",
	// 	}

	//   fmt.Println(getBook(adi))
	fmt.Println("=================empty interface===========")
	fmt.Println(ups(1))
	fmt.Println(ups(2))
	fmt.Println(ups(3))
	fmt.Println("=================empty interface===========")
	var adit map[string]string

	fmt.Println(adit == nil)

	fmt.Println("=================error interface===========")
	Division := func(value int, divider int) (int, error) {
		if divider == 0 {
			return 0, errors.New("cannot divided with zero")
		}
		return value / divider, nil
	}

	NumberDivided, err := Division(10, 0)

	if err == nil {
		fmt.Println(NumberDivided)
	} else {
		fmt.Println("error:", err)
	}

	fmt.Println("=================type assertion===========")
	var TryAssertion interface{} = ups(2)

	switch value := TryAssertion.(type) {
	case string:
		fmt.Println("string", value)
	case bool:
		fmt.Println("bool", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown")
	}

	fmt.Println("=================pointer===========")
	// golang is pass by value not pass reference
	// so if I pass varaible to other variable
	// it copy var value not variable reference
	// this make referenced var not changed if data from
	// other variable changed
	book1 := Book{"clean architecture", 300, "O'relly"}
	book2 := book1

	fmt.Println("example 1:")
	book2.BookName = "Clean architecture mark II"
	fmt.Println(book1) // value didn't change
	fmt.Println(book2)

	fmt.Println("example 2:")
	// adding pointer to reference variable
	book3 := &book1
	book3.BookName = "agustrioo"
	fmt.Println(book1) //value changed
	fmt.Println(book3)

	fmt.Println("example 3:")
	book4 := Book{"agus", 20, "will"}
	fmt.Println(book1) //value not change
	fmt.Println(book4)

	fmt.Println("example 4:")
	*book3 = Book{"agusyooo", 20, "will"}
	fmt.Println(book1) //value changed

	fmt.Println("example 5:")
	book5 := new(Book)
	book6 := book5

	book6.BookName = "yuudi"
	fmt.Println(book5)
	fmt.Println(book6)

	fmt.Println("=================pointer in func===========")
	changeBookName(book6)
	fmt.Println(book5)
}
