package main

import (
	"fmt"
)

//RANGE, CLOSE Y SELECT EN CHANNELS

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c)) //leng indica cuantos elementos hay,cap la capacidad maxima

	//range y close
	close(c) //cierra los channels una vez que dejas de usarlos

	//range sirve para iterar en cada uno de los elementos
	for message := range c {
		fmt.Println(message)
	}

	//select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}

//CHANNELS-forma de organizar routines,permiten la comunicacion y sincronizacion de datos en la gorutine

// func say(text string, c chan<- string) {
// 	c <- text //<- indica que vamos a ingresar un dato en ese canal
// }

// func main() {
// 	c := make(chan string, 1)

// 	fmt.Println("Hello")
// 	go say("Bye", c)

// 	fmt.Println(<-c) //debemos imprimir la salida de ese dato del canal,cdo lo coloco de lado izquiedo es salida
// }

//	GORUTINE
// es una forma de realizar concurrencia y ejecutar tareas de manera concurrente

// func say(text string, wg *sync.WaitGroup) {
// 	defer wg.Done() //hace que esta fn sea la ultima que se ejecute
// 	fmt.Println(text)
// }

// func main() { //esta fn esta corriendo en una gorutine xq una vez que termina la ejecucion muere
// 	var wg sync.WaitGroup //1 creamos un conjunto de gorutine y las va liberando poco a poco

// 	fmt.Println("Hello")
// 	wg.Add(1)
// 	go say("Word", &wg) //agregamos este keyword para que lo ejecute de forma concurrente(&wg es un puntero)

// wg.Wait()

// }
