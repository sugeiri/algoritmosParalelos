package main

//ARREGLOS Entero
import (
	"fmt"
	"os"
)
var (
	arreglo []int
)
func Menu() int {
	var x int
	fmt.Println("\nQue Desea Hacer?")
	fmt.Println("--------------------------------------------")
	fmt.Print("1-> Agregar Valor\n2-> Ver Todos los Valores \n3->Salir\n")
	fmt.Scan(&x)
	return x
}
func main()	{
	res := 0
	res= Menu()
	for ; res!=4; {
		  switch res {
	  case 1:
		lectura()
		break
	  case 2:
		fmt.Print("\n************************************\n")
		fmt.Println(arreglo)
		fmt.Print("\n************************************\n")
		break
	  case 3:
		os.Exit(3)
	  }
	  res= Menu()
	}
	os.Exit(3)
	  
}
func lectura() {
	var x int
	fmt.Println("\nIngrese Numero")
	fmt.Println("---------------------")
    fmt.Print("-> ")
	fmt.Scan(&x)
	arreglo=append(arreglo, x)
}

