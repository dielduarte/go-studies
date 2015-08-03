package main 

import "fmt"

func main(){
	//instancio e atribuio valor a variavel i
	i := 1

	//exemplo de um for simples em go
	// o mesmo segue uma ideia parecida de um while
	//enquanto i for menor ou igual a 3 faca...
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	//exemplo de um for completo em go
	for j := 7; j <= 9; j++ {
  	fmt.Println(j)
  }

  //for infinito em go
  for {
      fmt.Println("loop")
      break
  }




}