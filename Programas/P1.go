package main

//CALCULOS
import (
	"fmt"
)
var (
	a float32
	b float32

)
func Menu() int {
	var x int
	fmt.Println("\nQue Desea Hacer?")
	fmt.Println("--------------------------------------------")
	fmt.Print("1->Sumar dos Numeros\n2->Multiplicar dos Numeros \n3->Restar dos Numeros\n4->Dividir dos Numeros\n5->Salir\n")
	fmt.Scan(&x)
	return x
}
func main()	{
	res := 0
	res= Menu()
	for ; res!=5; {
		switch res {
			case 1:		
				//SUMAR
				lectura()
				x:=a+b
				fmt.Printf("\n**LA SUMA DE LOS DOS NUMEROS ES: %.0f**\n",x)

				break
			case 2:
				//MULTIPLICAR
				lectura()
				x:=a*b
				fmt.Printf("\n**LA MULTIPLICACION DE LOS DOS NUMEROS ES: %.0f**\n",x)
				break
			case 3:
				//RESTAR
				lectura()
				x:=a-b
				fmt.Printf("\n**LA RESTA DE LOS DOS NUMEROS ES: %.0f**\n",x)
				break
			case 4:
				//DIVIDIR
				lectura()
				x:=a/b
				fmt.Printf("\n**LA DIVISION DE LOS DOS NUMEROS ES: %.2f**\n",x)
				//LIMPIAR TODO
				break
			}
			res= Menu()
		}
}
func lectura(){
	fmt.Println("\nIngrese  El primer Numero")
	fmt.Println("---------------------")
    fmt.Print("-> ")
	fmt.Scan(&a)
	fmt.Println("\nIngrese  El Segundo Numero")
	fmt.Println("---------------------")
    fmt.Print("-> ")
	fmt.Scan(&b)
}
