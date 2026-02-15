package main

import (
	"fmt"
)

func main(){
	var s [] int

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

// func main(){
// 	s := make([] int, 3, 5)

// 	s[0] = 5
// 	s[2] = 10
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }

// func main(){
// 	s := make([] int, 3)
// 	s[0] = 5

// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }

// type User struct{
// 	Name string
// 	Age int
// 	Salary float64
// }

// func main(){
// 	user1 := User{
// 		Name: "Galib",
// 		Age: 26,
// 		Salary: 10,
// 	}

// 	p:= &user1

// 	fmt.Println(p.Name)
// }

// func print(numbers *[3]int){
// 	fmt.Println(numbers)
// }

// func  main()  {
// 	arr := [3] int {3, 6, 9}
// 	print(&arr)
// }

// var arr2 = [3] string{"I", "Love", "you"}

// func main (){
// 	arr := [2] int {3, 6}

// 	fmt.Println(arr[1])
// 	fmt.Println(arr2)
// }

// type User struct {
// 	Name string
// 	Age  int
// }

// func printUserDetails(usr User) {
// 	fmt.Println("Name: ", usr.Name)
// 	fmt.Println("Age: ", usr.Age)
// }

// func (usr User) userDetails(){
// 	fmt.Println("Name: ", usr.Name)
// 	fmt.Println("Age: ", usr.Age)
// }

// func (usr User) userInfo(a int){
// 	fmt.Println("Name: ", usr.Name)
// 	fmt.Println("Age: ", usr.Age)
// 	fmt.Println("New number: ", a)
// }

// func main() {
// 	user1 := User{ //instantiate
// 		Name: "Galib",
// 		Age:  26,
// 	}
// 	printUserDetails(user1)
// 	user1.userDetails()
// 	user1.userInfo(8)

// 	user2 := User{ //Instance or object
// 		Name: "Mamun",
// 		Age:  27,
// 	}
// 	printUserDetails(user2)
// 	user2.userDetails()
// 	user2.userInfo(10)
// }

// const a = 10

// var p = 100

// func call(){
// 	add:= func(x int, y int){
// 		z := x + y
// 		fmt.Println(z)
// 	}

// 	add(5, 6)
// 	add(p, a)
// }

// func main(){
// 	call()

// 	fmt.Println(a)
// }

// func init(){
// 	fmt.Println("Hello world")
// }

// func processOperation(a int, b int, op func(p int, q int)){
// 	op(a, b)
// }

// func add(a int, b int){
// 	c:= a + b
// 	fmt.Println(c)
// }

// func main(){
// 	processOperation(3, 5, add)
// }

// func main(){
// 	add := func(a int, b int){
// 		c := a + b
// 		fmt.Println(c)
// 	}

// 	add(2, 4)
// }

// func init(){
// 	fmt.Println("I'll be called first")
// }

// var a = 47

// func main() {
// 	var age = 80

// 	if age > 10 {
// 		a := 10
// 		fmt.Println(a)
// 	}
// 	fmt.Println(a)
// }

// var (
// 	a = 10
// 	b = 20
// )

// func printNum(num int) {
// 	fmt.Println(num)
// }

// func add(a int, b int){
// 	res := a + b
// 	printNum(res)
// }

// func main(){
// 	add(a, b)
// }

// import (
// 	"example.com/method"
// )

// var (
// 	a = 10
// 	b = 20
// )

// func main() {
// 	method.Add(20, 30)
// }

// func welcomeSection() {
// 	fmt.Println("Welcome to the application")
// }

// func getName() string {
// 	var name string
// 	fmt.Println("Please enter your name ?")
// 	fmt.Scanln(&name)
// 	return name
// }

// func getNumbers() (int, int) {
// 	var num1 int
// 	var num2 int

// 	fmt.Println("Enter your first number...")
// 	fmt.Scanln(&num1)
// 	fmt.Println("Enter your second number...")
// 	fmt.Scanln(&num2)
// 	return num1, num2
// }

// func addNumber(num1 int, num2 int) int {
// 	sum := num1 + num2
// 	return sum
// }

// func showDetails(name string, sum int) {
// 	fmt.Println("Hello", name)
// 	fmt.Println("Summation =", sum)
// }

// func lastMessage() {
// 	fmt.Println("Thank you for using the application")
// 	fmt.Println("Hav a good day.")
// }

// func main() {
// 	welcomeSection()
// 	name := getName()
// 	num1, num2 := getNumbers()
// 	sum := addNumber(num1, num2)
// 	showDetails(name, sum)
// 	lastMessage()
// }
