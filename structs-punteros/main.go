package main

import "fmt"

//INTERFACES- es un metodo donde puedes compartir otros metodos

type figura2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figura2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myReectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myReectangulo)

	//lista de interfaces-debe indicarse el tipo de dato que vamos a insertar
	myInterface := []interface{}{"Hola", 12, 13.4}
	fmt.Println(myInterface...)

}

//STRINGERS:PERSONALIZAR EL OUTPUT DE STRUCTS

// type pc struct {
// 	ram   int
// 	brand string
// 	disk  int
// }

// // para personalizar el autput de consola-sprint crea un string
// func (myPC pc) String() string {
// 	return fmt.Sprintf("Tengo %d GB RAM, %d  GB de disco y es %s", myPC.ram, myPC.disk, myPC.brand)
// }

// func main() {
// 	myPC := pc{ram: 16, brand: "msi", disk: 100} //instanciar
// 	fmt.Println(myPC)
// }

//struct y punteros
// type pc struct {
// 	ram   int
// 	disk  int
// 	brand string
// }

// func (myPC pc) ping() {
// 	fmt.Println(myPC.brand, "Pong")
// }

// DUPLICA RAMA
// func (myPC *pc) duplicateRAM() {
// 	myPC.ram = myPC.ram * 2 //duplicamos la RAM
// }

// func main() {
// 	a := 50
// 	b := &a //apunta a la direccion en la memoria

// fmt.Println(b)
// fmt.Println(*b) //el * accede al valor de esa direccion de memoria

//apuntar a la misma direccion de memoria
// *b = 100
// fmt.Println(a)

// myPC := pc{ram: 16, disk: 200, brand: "msi"}
// fmt.Println(myPC)

// myPC.ping()

// //ram nos da 16
// fmt.Println(myPC)
// myPC.duplicateRAM()

//duplica y la RAM da 32
// fmt.Println(myPC)
// myPC.duplicateRAM()

//duplica RAM da 64
// fmt.Println(myPC)

//}

//los punteros son accesos a la memoria
//variable > direccion de memoria > valor de variable
