package main 

import (
	"fmt"
	"math"
)

//declaracao de uma constant do tipo stirng
const s string = "constant"

func main(){
	fmt.Println(s)

	//constant pode ser declarada sem tipo e pega o tipo 
	//do seu primeiro valor atribuido
	const n = 500000000

	
  const d = 3e20 / n
    
  fmt.Println(d)

  fmt.Println(int64(d))

  fmt.Println(math.Sin(n))
}