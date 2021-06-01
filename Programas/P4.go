package main

//SUMA DE VALORES EN ARREGLO
import (
	"fmt"
)
var (
	arreglo []int
)
func main()	{
	lectura()
	Sumar()
	  
}
func lectura() {
	var x int
	x=1
	fmt.Println("\nIngrese los Valores (Ingrese 0 para terminar)")
	fmt.Println("-------------------------------------------")
	for ;x!=0;
	{
		fmt.Print("-> ")
		fmt.Scan(&x)
		if x!=0{
			arreglo=append(arreglo, x)
		}
	}
}
func Sumar() {
	var x,s int
	for x=0;x<len(arreglo);x++{
		s+=arreglo[x]
	}
	fmt.Printf("\n LA SUMA DE LOS NUMEROS INGRESADOS ES: %d",s)
}

