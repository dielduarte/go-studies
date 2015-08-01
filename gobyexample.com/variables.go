package main 

import "fmt"

func main() {
    var a string = "variavel tipada em  go, nesse caso string"
    fmt.Println(a)

    //multiplos valores em go
    var b, c int = 1, 2
    fmt.Println(b, c)

    //variavel pode ser inicializada sem typagem declarada,
    //porem a mesma recebe o tipo inicial e nao pode receber 
    //outro tipo no decorrer do programa
    var d = true
    fmt.Println(d)

    //variavel tbm pode ser declarada sem valor inicial
    //quando isso acontece a mesma recebe dinamicamente um valor default 
    //para o seu tipo no caso do exemplo (int)
    var e int
    fmt.Println(e)

    //a sintaxy := e usada para declarar e atribuir valor 
    //a uma variavel que ainda nao foi declarada
    f := "string"
    fmt.Println(f)  
}