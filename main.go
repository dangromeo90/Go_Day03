package main

import (
	"fmt"
	"math"
)
func main()  {
	addEmployee(Employee{id : 1 , name : "Nguyễn Vũ Nhật Đăng" , age : 34 , position: "Student"})
	addEmployee(Employee{id : 2 , name : "Lê Minh Tân" , age : 25 , position: "Student"})
	addEmployee(Employee{id : 3 , name : "Võ Nguyễn Thiên Ân" , age : 29 , position: "Student"})
	showMaxAge()
	//Chu vi , diện tích hình tròn
	line()
	tinh()
	// Chu vi , diện tích hình tròn và tam giác (ỉnterface)
	line()
	CircleRectangle()
	line()
	ShowVehicle()
	line()
	swapAB()
	line()
	increase()
	line()
	Factorial()
	line()
	Fibonacci()

}

// Câu 1 Quản lý thông tin nhân viên
type Employee struct{
	id int
	name string
	age int
	position string
}

var listEmployee = make(map[int]Employee)

func addEmployee(emp Employee){
	listEmployee[emp.id] = emp
	fmt.Println("Thêm nhân viên thành công",emp.name)
}

func showMaxAge()  {
	fmt.Println("Danh Sách Nhân Viên có tuổi cao nhất")
	if len(listEmployee)==0 {
		fmt.Println("Danh sách trống")
	}

	maxAge := -1
	var maxEmp Employee
	for _, emp := range listEmployee{
		if emp.age > maxAge{
			maxAge = emp.age
			maxEmp = emp
		}
	}
	fmt.Printf("ID: %d, Tên: %s, Tuổi: %d, Vị trí: %s\n", maxEmp.id, maxEmp.name, maxEmp.age, maxEmp.position)
}
//Câu 2 : Tính chu vi và diện tích hình tròn

type Circle struct{
	radius float64
}
func (c Circle) chuVi() float64  {
	return 2 * math.Pi * c.radius 
}
func (c Circle) dientich() float64  {
	return math.Pi * c.radius * c.radius
}
func tinh()  {
	circle := Circle{radius: 5}
	fmt.Printf("Chu vi: %.2f\n", circle.chuVi())
	fmt.Printf("Dien tich: %.2f\n", circle.dientich())
}

// Bài tập vè Interface
// Câu 1 
type Shape interface {
	Area() float64
	Perimeter() float64
}
func (c Circle) Area() float64  {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.radius
}

type Rectangle struct{
	width, height float64
}
func (r Rectangle) Area() float64{
	return r.width * r.height
}
func (r Rectangle) Perimeter()float64  {
	return (r.height + r.width)*2
}
func CircleRectangle()  {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4 ,height: 6}
	fmt.Printf("Chu vi hình tròn: %.2f\n" , circle.Perimeter())
	fmt.Printf("Diện tích hình tròn: %.2f\n" , circle.Area())
	fmt.Printf("Chu vi tam giác: %.2f\n" , rectangle.Perimeter())
	fmt.Printf("Diện tích tam giác: %.2f\n" , rectangle.Area())

}
 // Câu 2 - Verhicle

	type Vehicle interface{
		Start() string
		Stop() string
	}
	type Car struct{
		name string
	}
	type Bike struct {
		name string
	}
	func (c Car) Start() string {
		return c.name + " is starting"
	}
	func (c Car) Stop() string {
		return c.name + " is stopping"
	}
	func (b Bike) Start() string {
		return b.name + " is starting"
	}
	func (b Bike) Stop() string {
		return b.name + " is stopping"
	}
	func ShowVehicle()  {
		car := Car{name: "Tesla model 3"}
		bike := Bike{name : "Yamaha"}
		vehicles := []Vehicle{car , bike}
		for _, vehicle := range vehicles{
			fmt.Println(vehicle.Start())
			fmt.Println(vehicle.Stop())
		}
	}

// Con trỏ
// Câu 1
func ChuyenDoi(a,b * int)  {
	tam := *a
	*a = *b
	*b = tam
}
func swapAB()  {
	a := 3
	b := 5
	fmt.Printf("a = %d , b = %d \n", a,b) 
	ChuyenDoi(&a,&b)
	fmt.Printf("Chuyển đổi : a = %d, b = %d \n" ,a , b )
}
// Câu 2
func tang(x * int)  {
	
	*x = *x + 1
}
func increase(){
	x := 10 
	fmt.Printf("x = %d\n", x)
	tang(&x)
	fmt.Printf("x tăng 1 = %d\n",x)
}

// Đệ quy
// Câu 1
func giaithua(n int) int  {
	s := 1
	for i := 1 ; i <= n ; i++{
		s *= i
	}
	return s
}
func Factorial()  {
	n := 5 
	fmt.Printf("Giai thừa %d = %d", n, giaithua(n))
}

// Câu 2
func fibo(n int) int  {
	if(n == 0){
		return 0
	}
	if(n == 1){
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
func PrintFibo(n int)  {
	arr := make([]int , n )
	for i := 0 ; i < n ; i++{
		arr[i] = fibo(i)
	}
	fmt.Printf("Dãy Fibonacci của %d: {", n)
	for i, num := range arr {
		if i == len(arr)-1{
			fmt.Printf("%d",num)
		}else{
			fmt.Printf("%d, ", num)
		}
	}
	fmt.Println("}")
}
func Fibonacci()  {
	n := 6
	PrintFibo(n)
}


func line()  {
	fmt.Println("----------------------------------------")
}