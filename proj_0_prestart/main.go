package main

import (
	"fmt"

	// "proj_1/configs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("first")

	app := fiber.New()

	//* package
	// MyHelperFuction()
	// configs.LoadAppSettings()
	//---

	//* types
	//Basic Types: int, float64, string , boolean
	//Composite Types: array, slice, map, struct
	//Pointer types: *

	// var age int
	// var height float64
	// var firstName string
	// var isEmployed bool
	// age = 21
	// height = 138.5
	// firstName = "Joe"
	// isEmployed = true
	// fmt.Println(age, height, firstName, isEmployed)

	//* printf
	// fmt.Printf("Age: %d", age)
	// fmt.Printf("Height: %f", height)
	// fmt.Printf("FirstName: %s", firstName)
	// fmt.Printf("IsEmployed: %v", isEmployed)

	// if age > 65 {
	// 	fmt.Println("Senior citizen")
	// } else if age > 16 {
	// 	fmt.Println("audult")
	// }

	// for i := 0; i < 10; i++ {
	// }

	//* array
	// var myFamily [3]string
	// myFamily[0] = "Jay"
	// myFamily[1] = "jiane"
	// myFamily[2] = "carol"
	// fmt.Println("My family: %v", myFamily)

	// myFamily := [3]string{"jay", "jane", "carol"}
	// myFamily[2] = "kate"
	// fmt.Println("My family: %v", myFamily)

	// myCourse := [3][2]string{
	// 	{"go", "nodejs"},
	// 	{"aws", "gcp"},
	// 	{"cdk", "pulumi"},
	// }
	// fmt.Println("Available Course %v", myCourse)

	//* slice
	// var myFriends []string
	// myFriends = append(myFriends, "Mike", "Adam")
	// myFriends = append(myFriends, "Scola", "Sam")
	// fmt.Println("My friends: %v", myFriends)

	// mySliceCourse := [][]string{
	// 	{"go", "nodejs"},
	// 	{"aws", "gcp"},
	// 	{"cdk", "pulumi"},
	// }
	// course := []string{"IAC", "Cloud formation"}
	// mySliceCourse = append(mySliceCourse, course)
	// fmt.Println("Available Course %v", mySliceCourse)

	//* map
	// myWishlist := make(map[string]string)
	// myWishlist["first"] = "MacPro"
	// myWishlist["second"] = "900 Billion Dollar"
	// myWishlist["third"] = "a beautiful car"
	// delete(myWishlist, "third")
	// firstWish := myWishlist["first"]
	// log.Println(firstWish)
	// fmt.Printf("My wish lish %v", myWishlist)

	//* struct
	// type Details struct {
	// 	Description string `json:"description"`
	// 	Images      string `json:"images"`
	// }
	// type Product struct {
	// 	Name    string  `json:"name"`
	// 	Price   float64 `json:"price"`
	// 	Details `json:"details"`
	// }
	// // var product Product
	// product := Product{
	// 	Name:  "MacPro",
	// 	Price: 9000,
	// 	Details: Details{
	// 		Description: "An incredible machine",
	// 		Images:      "http://macrpimage.jpg",
	// 	},
	// }
	// product.Name = "ABC macbook"
	// fmt.Println("Product struct: %v", product)

	//* conditional statements
	// pointers

	// if else
	// age := 29
	// if age > 65 {
	// 	fmt.Println("Senior Citizen")
	// } else if age > 17 {
	// 	fmt.Println("Adult Citizen")
	// } else {
	// 	fmt.Println("Child")
	// }

	// switch case
	// seatClass := "FirstClass"
	// switch seatClass {
	// case "FirstClass":
	// 	fmt.Println("you will get free drinks")
	// case "BussinessClass":
	// 	fmt.Println("You will get more legrooms")
	// default:
	// 	fmt.Println("You need to pay for services")
	// }

	// select
	// var myFriends []string
	// for i := 0; i < 10; i++ {
	// 	myNewFriends := fmt.Sprintf("Friend %d", i)
	// 	myFriends = append(myFriends, myNewFriends)
	// }
	// fmt.Println(myFriends)

	// for index, value := range myFriends {
	// 	fmt.Println(index, value)
	// }

	// isOver := 0
	// for {
	// 	isOver++
	// 	if isOver > 99 {
	// 		return
	// 	}
	// }

	//* Ponters
	// jay := "laptop"
	// fmt.Println(jay)
	// fmt.Println(&jay)	//address
	// var guest *string	//pointer <nil>
	// guest = &jay			//address
	// fmt.Println(guest)	//address
	// fmt.Println(*guest) //value

	//* function -> return
	//Nhóm In ra màn hình (fmt.Print và print): Dùng khi bạn muốn người dùng hoặc lập trình viên nhìn thấy thông báo ngay trên console.

	//Nhóm Xử lý chuỗi (fmt.Sprint): Dùng khi bạn không muốn in ngay, mà muốn gom các giá trị lại thành một đoạn văn bản (string) để:
	// Ghi vào file log.
	// Gửi dữ liệu qua mạng (HTTP response, WebSocket).
	// sayHello()
	// getUserName()
	// fmt.Printf("returning from fuction: %v", getUserName())

	// name, age := getUserById(1, 16)
	// fmt.Printf("Name: %v, Age: %v", name, age)

	// fmt.Printf("Total Amount: %v", calculateTotal(1.2, 4.5, 6.6))

	// concateUserName := func(fname string, lname string) string {
	// 	return fmt.Sprintf("%s, %s", fname, lname)
	// }

	// fmt.Printf("user full name is: %s", concateUserName("jav", "max"))

	//* Receiver function
	// p := Product{ //Composite Literal (hay cách khởi tạo giá trị cho Struct)
	// 	Name:  "MacPro",
	// 	Price: 8000,
	// 	Stock: 1,
	// }

	// fmt.Printf("Total Amount %f", p.Calculate(2))

	// p.reduceStock(1)
	// fmt.Println(p)

	app.Listen("localhost:9000")
}

// type Product struct {
// 	Name  string
// 	Price float64
// 	Stock int
// }

// func (p Product) Calculate(qty int) float64 {
// 	return p.Price * float64(qty)
// }

// func (p Product) reduceStock(qty int) {
// 	if p.Stock >= qty {
// 		p.Stock -= qty
// 	}
// }

// func sayHello() {
// 	fmt.Println("hello jay")
// }

// func getUserName() string {
// 	return "my name jay"
// }

// func getUserById(id int, ageInput int) (name string, age int) {
// 	//db call
// 	println(id, ageInput)
// 	return "jay", 36
// }

// // Variadic Parameter (Tham số biến thiên).
// // /Dấu ...: Cho phép hàm calculateTotal nhận vào một hoặc nhiều giá trị kiểu float64. Bạn không bị giới hạn số lượng tham số truyền vào khi gọi hàm.
// func calculateTotal(products ...float64) float64 {
// 	totalAmount := 0.0
// 	for _, price := range products {
// 		totalAmount += price
// 	}
// 	return totalAmount
// }
