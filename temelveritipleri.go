package main

import "fmt"

func main() {

	var name string = "Emir"
	var yas int = 23
	var evlimi bool = true
	var weight float32 = 85.5

	fmt.Println(name)
	fmt.Println(yas)
	fmt.Println(evlimi)
	fmt.Println(weight)

	//Veri Tipi Sorgulama

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", yas)
	fmt.Printf("%T\n", evlimi)

	// Kısa Şekilde Veri Tanımlama

	var ad = "Emir"
	ad2 := "Emir2"

	fmt.Println(ad)
	fmt.Println(ad2)

	// Tek Var İçerisine Tanımlama

	var (
		name2   string  = "Emir"
		yas2    int     = 23
		evlimi2 bool    = true
		weight2 float32 = 85.5
	)
	fmt.Println(name2)
	fmt.Println(yas2)
	fmt.Println(evlimi2)
	fmt.Println(weight2)

	//------------------------------------

	var name3, yas3, evlimi3, weight3 = "Emir3", 40, true, 85.5
	fmt.Println(name3)
	fmt.Println(yas3)
	fmt.Println(evlimi3)
	fmt.Println(weight3)
}

/* BASIC VERİ TİPLERİ
NUMERIC TYPES
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

STRING TYPES
A string type represents the set of string values.

BOOLEAN TYPES
A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.

*/
