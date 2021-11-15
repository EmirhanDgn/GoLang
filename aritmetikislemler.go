package main

import "fmt"

func main() {

	x, y := 15, 20

	fmt.Printf("%T, %v\n", (x + y), (x + y))
	fmt.Printf("%T, %v\n", (x / y), (x / y))

	z := 5.0 / 2 // Virgüllü, float bir sonuç almak için değerlerden birisi küsüratlı gösterilmelidir.
	fmt.Printf("%T, %v\n", z, z)

	a := 10
	fmt.Println(a)
	fmt.Println(a + 1)
	fmt.Print(a - 1)

}
