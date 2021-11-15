package main

import "fmt"

var PackageVar = "Kod Satırının Her Yerinden Ulaşılabilen Veri"

func main() {

	var name = "Emir"
	name, surname := "Ahmet", "Doğan"
	fmt.Println(name, surname)

	/*
		var funcVar = "Func Scope"

		fmt.Println(funcVar)
		fmt.Println(PackageVar)
	*/
}
