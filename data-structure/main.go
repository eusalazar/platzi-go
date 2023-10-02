package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

}

//Llave valor con maps
//m := make(map[string]int)
//int es el tipo de dato-guardamos las edades de ciertas personas
//el 0 value de int es 0
//m["Jose"] = 14
//m["Pepito"] = 20

//fmt.Println(m)

//recorrer un par, recorrer su llave y su valor
//map se ejecuta de forma aleatoria. slice los retorna en el mismo orden
// for i, v := range m {
// 	fmt.Println(i, v)
// }

//encontrar el valor-el ok nos indica si esta llave que estamos buscando existe
//cuando no tenemos certeza si una llave esta en su diccionario

// value, ok := m["Josep"]
// fmt.Println(value, ok)

// Esta funcion detecta las palabras de tipo palindromo
//func isPalindromo(text string) {
//strings.ToLower se utiliza para convertir text a miniscula
// text = strings.ToLower(text)
// var textReverse string

// for i := len(text) - 1; i >= 0; i-- {
// 	textReverse += string(text[i])
// }
// if text == textReverse {
// 	fmt.Println("Es palindromo")
// } else {
// 	fmt.Println("No es un palindromo")
// }
//}

// func main() {
// 	texto := "Anilina"
// 	isPalindromo(texto)

// 	isPalindromo("Amor A roma")

//Recorrido con slice
//slice := []string{"hola", "que", "tal"}

// for i := range slice { //sino necesitamos el indice escapamos _
// 	fmt.Println(i) //si necesitamos solo el indice solo usamos i
// }

//Array- son inmutables
//var array [4]int // valores inmutables
// array[0] = 1
// array[1] = 2
//len indica cuantos elementos hay, cap indica la capacidad max
// fmt.Println(len(array), cap(array))

//slice,son mutables, no se le indica la cantidad de valores que va a tener
// slice := []int{0, 1, 2, 3, 4, 5, 6}
// fmt.Println(slice, len(slice), cap(slice))

//metodos en el slice-es lo que se utiliza para trabajar con slice, array
//para interactuar con c/u de sus elementos

// fmt.Println(slice[0])
// fmt.Println(slice[:3])  //hasta,el nro que va desp de los puntos es exclusivo
// fmt.Println(slice[2:4]) //el primer numero es inclusivo
// fmt.Println(slice[4:])

//Append-como adicionar elemtos al slice
// slice = append(slice, 7)
// fmt.Println(slice)

//Append-como adicionar listas al slice
// newSlice := []int{8, 9, 10}
// slice = append(slice, newSlice...) //descomprime este nuevo slice y en vez de agregar la lista completa agrega los elementos de forms indep
// fmt.Println(slice)
