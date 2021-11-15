package main

import "fmt"

func main() {
	/*
		var x, y, sum int

		x = 5
		y = 10
		sum = x + y
		fmt.Printf("%d ve %d Toplamı %d", x, y, sum)
	*/

	fmt.Println(sum(5, 10))

	merhaba() // Fonskiyonu Çalıştırma

	sum2(6, 15)
}

//func yazdıktan sonra fonksiyon ismi yazılmalıdır.
//func funcName(parameters) return type { code }
func sum(x, y int) int {
	return x + y
}

func merhaba() {
	fmt.Println("Fonksiyon Denemesi")
}

func sum2(x, y int) {
	fmt.Println(x + y)
}
