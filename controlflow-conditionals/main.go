package main

import "fmt"

func main() {
	// Defer(It is the keyword that executes the function before everything dies)
	// Used to close connections to databases or files after use.
	defer fmt.Println("Hola")
	fmt.Println("mundo")

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue It is used when a condition given in the for is interesting to continue
		//err
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break when a certain condition is met and you want the code to finish executing
		if i == 8 {
			fmt.Println("break")
			break
		}
	}

}

// nested conditionals
//We use switch when we have to use multiple nested conditionals
//even or odd
// modulo := 5 % 2
// switch modulo {
// switch modulo := 4 % 2; modulo {
// case 0:
// 	fmt.Println("Es par")
// default:
// 	fmt.Println("Es impar")
// }

//Without condition(sin)nested-anidado
// value := 200
// switch {
// case value > 100:
// 	fmt.Println("Es mayor a 100")
// case value < 0:
// 	fmt.Println("Es menor a 0")
// default:
// 	fmt.Println("No condicion")
// }

//conditional

// value1 := 1
// value2 := 2ß

// if value1 == 1 {
// 	fmt.Println("it is  1")
// } else {
// 	fmt.Println("not it 1")
// }

// //with and- Both conditions must be true for the code to run.
// if value1 == 1 && value2 == 3 {
// 	fmt.Println("it is true")
// }

// //with or
// if value1 == 0 || value2 == 2 {
// 	fmt.Println("it is true, OR")
// }

// //convert text to number
// value, err := strconv.Atoi("53")
// if err != nil {
// 	log.Fatal(err) //If there is an error, the code ends.
// }
// fmt.Println("value:", value)

//In go there is only one type of iterative FOR loop
//for conditional
// for i := 0; i < 10; i++ {
// 	fmt.Println(i)
// }

//For while(It is a stop until a condition is met)
// counter := 0
// for counter < 10 {
// 	fmt.Println(counter)
// 	counter++ // If we don't tell it to increase the counter it will stay at 0

// }

//For forever it does not have an end
// counterForever := 0
// for {
// 	fmt.Println(counterForever)
// 	counterForever++
// }
