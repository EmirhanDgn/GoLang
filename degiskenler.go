package main

import (
	"fmt"
	"reflect"
)

func main() {

	//String Değişkeni Ataması
	var a string = "String Değişkeni A"

	var b string = "String Değişkeni B"

	fmt.Println(a)
	fmt.Println(b)

	//Aynı Anda Birden Çok Değişken Tanımlama
	var adi, soyadi string = "Emir", "Doğan"

	fmt.Println(adi, soyadi)

	//İnt Değişkeni Ataması
	var i int = 5
	fmt.Println(i)

	//Tek Satırda Çift Tanımlama
	var k, m int = 2, 3

	fmt.Println(k)
	fmt.Println(m)

	//Reflect komutu ile değişken türünü görebiliyoruz.
	c := "portkal"
	fmt.Println(reflect.TypeOf(c))

	// Bool Değişkeni
	var isMarried bool
	isMarried = true
	fmt.Println(isMarried)

}
