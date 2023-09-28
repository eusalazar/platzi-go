package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int { // es el valor de salida
	return a * 2
}

// si necesito que retorne 2 o mas valores
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("value:", value)

	//la forma de retornar 2 valores
	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2:", value1, value2)

	// value1, _ := doubleReturn(2)   si solo necesito un solo valor
	// fmt.Println("value1:", value1,)

}

// declaracion de const(estatucamente tipado)
// const pi float64 = 3.14
// const pi2 = 3.164
// fmt.Println("pi:", pi)
// fmt.Println("pi2:", pi2)

//Declaracion de variables enteras
// base := 12
// var altura int = 14
// var area int
// fmt.Println(base, altura, area)

//zero values es un valor de que va a tener por defecto
//una variable que no le asignamos un valor
//en otros lenguajes el valor en null
// var a int
// var b float64
// var c string
// var d bool

// fmt.Println(a, b, c, d)

//area de un cuadrado
// const baseCuadrado = 10
// areaCuadrado := baseCuadrado * baseCuadrado
// fmt.Println("areaCuadrado:", areaCuadrado)

//operadores aritmeticos
// x := 10
// y := 50

//suma
// result := x + y
// fmt.Println("suma:", result)

//resta
// result = y - x
// fmt.Println("resta:", result)

//multiplicacion
// result = x * y
// fmt.Println("multipliacion:", result)

//division
// result = y / x
// fmt.Println("division:", result)

//modulo(par o impar)
// result = y % x
// fmt.Println("modulo:", result)

//incrementar/anadir 1, en un principio tenemos denifinada x como 10
// x++
// fmt.Println("incrementar:", x)

//decrementar(resta uno)
// x--
// fmt.Println("decrementar:", x)

//area de un rectangulo
// const largo = 2
// const ancho = 3
// areaRectangulo := largo * ancho
// fmt.Println("areaRectangulo:", areaRectangulo)

//area de un trapecio
// const baseMayor = 8
// const baseMenor = 4
// const alturaTrapecio = 6

// areaTrapecio := (((baseMayor + baseMenor) / 2) * alturaTrapecio)
// fmt.Println("areaTrapecio:", areaTrapecio)

//paquete fmt es el paquete con el que podemos interactuar en consola
// helloMesagge := "Hello"
// worldMessage := "world"

//Println/ salto de linea
// fmt.Println(helloMesagge, worldMessage)
// fmt.Println(helloMesagge, worldMessage)

//Print imprime y agrega una funcion extra
//%s le decimos que recibe un string
//%d le decimos que recibe un entero
//%v cuando no sabemos que tipo de dato vamos a recibir
// nombre := "platzi"
// cursos := 500

// fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

// Sprintf no imprime nada en consola pero lo guarda como string
// message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
// fmt.Println(message)

//Nos dice el tipo de datos que tenemos
// fmt.Printf("helloMesagge:%T \n", helloMesagge)
// fmt.Printf("cursos: %T\n", cursos)
