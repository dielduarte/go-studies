package main

import "fmt"

func main() {

		//teste se 7 % 2 e impar ou par
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    //testa de 8 e divisivel por 4
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    //atribui valor a num e testa se o mesmo 
    //e negativo
    // ou tem apenas um digito
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

}