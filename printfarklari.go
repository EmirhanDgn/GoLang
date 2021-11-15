package main

import "fmt"

func main() {

	name := "Emir"
	/*
		fmt.Print(name)
		fmt.Printf(name)
		fmt.Println(name)
	*/

	fmt.Print("Benim Adım", name) // Metinler Arası Boşluk Bırakmıyor ve Alt Satıra Atmıyor.
	fmt.Println("")
	fmt.Println("Benim Adım", name) // Verilen ifadeleri tam anlamıyla ortaya çıkarmaktadır.
	fmt.Printf("Benim Adım", name)  // Verilen string ifadeyi yazdırır fakat verilen değişken ise onu formatı halinde yazdırır.
}
